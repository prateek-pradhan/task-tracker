# 📝 Task Tracker CLI

A simple **command-line task tracker** built in **Go**, powered by [Cobra](https://github.com/spf13/cobra).  
It allows you to add, update, delete, and list tasks, persisting them in a local `tasks.json` file.  

---

## 🚀 Features
- Add, update, delete tasks
- Update task status (`To Do`, `In Progress`, `Done`)
- List tasks with optional filters
- Stores tasks in `tasks.json`

---

## ⚙️ Installation

```bash
git clone https://github.com/prateek-pradhan/task-tracker.git
cd task-tracker
go build -o task-tracker-cli main.go
mv task-tracker-cli /usr/local/bin/   # optional
```

---

## 📌 Usage

```bash
task-tracker-cli [command]
```

### Commands

- **Add**: `task-tracker-cli add "Buy fruits"`  
- **Update description**: `task-tracker-cli update 1 "Buy vegetables"`  
- **Update status**: `task-tracker-cli mark-in-progress 1` / `mark-done 1`  
- **Delete**: `task-tracker-cli delete 1`  
- **Get**: `task-tracker-cli get 1`  
- **List**: `task-tracker-cli list` or filter by `todo`, `in-progress`, `done`  

---

## 🗂 Data Storage

Tasks are stored in `tasks.json` like this:

```json
{
  "Tasks": {
    "1": {
      "Description": "Buy fruits",
      "Status": "To Do",
      "CreatedAt": "2025-08-29T12:34:56Z",
      "UpdateAt": "2025-08-29T12:34:56Z"
    }
  },
  "NextId": 1
}
```

---

## 🛠 Notes
- Uses Go’s `encoding/json` for persistence  
- Auto-creates `tasks.json` if missing  
- Extendable with Cobra for more commands  
