package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Description string
	Status string
	CreatedAt time.Time
	UpdateAt time.Time
}

type TaskTracker struct {
	Tasks map[int]Task
	NextId int
}


func addTask(task string) (result string) {
	if len(task) == 0 {
		fmt.Println("Empty string no task has been added")
		return
	}

	tasks, err := readJson()
	if err != nil {
		return
	}

	nextId := tasks.NextId + 1
	tasks.NextId = nextId

	var taskObj Task
	taskObj.Description = task
	taskObj.Status = "To Do"
	taskObj.CreatedAt = time.Now()
	taskObj.UpdateAt = time.Now()
	
	if tasks.Tasks == nil {
    tasks.Tasks = make(map[int]Task)
	}

	tasks.Tasks[nextId] = taskObj
	writeJson(tasks)
	return fmt.Sprintf("Task added with id %d", nextId)
}

func updateTask(id string, task string) (result string) {
	if len(task) == 0 {
		fmt.Println("Empty string no task has been added")
		return
	}

	taskId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error: Invalid task id")
		return
	}

	tasks, err := readJson()
	if err != nil {
		return 
	}

	if _, exists := tasks.Tasks[int(taskId)]; !exists {
    return fmt.Sprintf("Error: Task %d does not exist", taskId)
	}

	if int(taskId) > tasks.NextId {
		fmt.Println("Error: Invalid task id")
		return
	}
	var taskObj Task
	taskObj.Description = task
	taskObj.Status = tasks.Tasks[int(taskId)].Status
	taskObj.CreatedAt = tasks.Tasks[int(taskId)].CreatedAt
	taskObj.UpdateAt = time.Now()
	
	tasks.Tasks[int(taskId)] = taskObj
	writeJson(tasks)
	return fmt.Sprintf("Task updated with id %d", taskId)
}

func updateTaskStatus(id string, newStatus string) (result string) {

	taskId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error: Invalid task id")
		return
	}

	tasks, err := readJson()
	if err != nil {
		return 
	}

	if _, exists := tasks.Tasks[int(taskId)]; !exists {
    return fmt.Sprintf("Error: Task %d does not exist", taskId)
	}

	if int(taskId) > tasks.NextId {
		fmt.Println("Error: Invalid task id")
		return
	}
	var taskObj Task
	taskObj.Description = tasks.Tasks[int(taskId)].Description
	taskObj.Status = newStatus
	taskObj.CreatedAt = tasks.Tasks[int(taskId)].CreatedAt
	taskObj.UpdateAt = time.Now()
	
	tasks.Tasks[int(taskId)] = taskObj
	writeJson(tasks)
	return fmt.Sprintf("Task updated with id %d", taskId)
}

func deleteTask(id string) (result string) {

	taskId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error: Invalid task id")
		return
	}

	tasks, err := readJson()
	if err != nil {
		return 
	}

	if _, exists := tasks.Tasks[int(taskId)]; !exists {
    return fmt.Sprintf("Error: Task %d does not exist", taskId)
	}

	if int(taskId) > tasks.NextId {
		fmt.Println("Error: Invalid task id")
		return
	}

	delete(tasks.Tasks, int(taskId))
	
	writeJson(tasks)
	return fmt.Sprintf("Task deleted with id %d", taskId)
}

func getTask(id string) (result string) {

	taskId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error: Invalid task id")
		return
	}

	tasks, err := readJson()
	if err != nil {
		return 
	}

	if _, exists := tasks.Tasks[int(taskId)]; !exists {
    return fmt.Sprintf("Error: Task %d does not exist", taskId)
	}

	if int(taskId) > tasks.NextId {
		fmt.Println("Error: Invalid task id")
		return
	}

	writeJson(tasks)
	return fmt.Sprintf("Task with id %d: %s", taskId, tasks.Tasks[int(taskId)])
}

func list(status string) {
	tasks, err := readJson()
	if err != nil {
		return 
	}

	var filterStatus []string
	switch status {
	case "":
		filterStatus = []string{"To Do", "In Progress", "Done"}
		break
	case "todo":
		filterStatus = []string{"To Do"}
		break
	case "in-progress":
		filterStatus = []string{"In Progress"}
		break
	case "done":
		filterStatus = []string{"Done"}
		break
	default:
		fmt.Println("Error: Invalid status filter. Use one of: todo, in-progress, done")
		return
	}

	for id, task := range tasks.Tasks {
		for _, s := range filterStatus {
			if task.Status == s {
				fmt.Printf("ID: %d, Description: %s, Status: %s\n", id, task.Description, task.Status)
				break
			}
		}
	} 
}

func readJson() (TaskTracker, error) {
    file, err := os.Open("./tasks.json")
    if os.IsNotExist(err) {
        return TaskTracker{Tasks: make(map[int]Task), NextId: 0}, nil
    }
    if err != nil {
        return TaskTracker{}, fmt.Errorf("unable to fetch tasks.json: %v", err)
    }
    defer file.Close()

    var tasks TaskTracker
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&tasks); err != nil {
        if err == io.EOF {
            return TaskTracker{Tasks: make(map[int]Task), NextId: 0}, nil
        }
        return TaskTracker{}, fmt.Errorf("error decoding JSON: %v", err)
    }

    // Ensure map is always initialized
    if tasks.Tasks == nil {
        tasks.Tasks = make(map[int]Task)
    }

    return tasks, nil
}


func writeJson(tasks TaskTracker) error {
    file, err := os.Create("./tasks.json")
    if err != nil {
        return fmt.Errorf("unable to open tasks.json: %v", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ") // optional: pretty print
    if err := encoder.Encode(tasks); err != nil {
        return fmt.Errorf("error encoding JSON: %v", err)
    }

    return nil
}
