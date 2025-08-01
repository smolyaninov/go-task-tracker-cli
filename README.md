# Task CLI

This is my first project in Go, completed as part of the [roadmap.sh Task Tracker](https://roadmap.sh/projects/task-tracker) challenge.

The goal was to learn the basic syntax and idioms of Go through hands-on practice, and to build a minimal but cleanly structured CLI application for task management (To‑Do tracker).

## Features

- Add new tasks
- Update task descriptions
- Delete tasks
- Mark tasks as in-progress or done
- List all tasks, or filter by status (`todo`, `in-progress`, `done`)
- Persistent JSON storage in the current directory

## Usage

```bash
task-cli add "Write unit tests"
task-cli list
task-cli list done
task-cli update 1 "Write integration tests"
task-cli mark-in-progress 1
task-cli mark-done 1
task-cli delete 1
```

## Build

```bash
go build -o task-cli ./cmd/task-cli
```
