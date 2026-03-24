# Task Tracker CLI

A simple and efficient command-line interface (CLI) for managing your daily tasks. Built with Go and the [Cobra](https://github.com/spf13/cobra) library.

## Features

- **Add Tasks**: Quickly create new tasks with a description and status.
- **Update Tasks**: Modify existing task descriptions or update their status.
- **Delete Tasks**: Remove tasks by their ID.
- **List Tasks**: View all tasks or filter them by status (`todo`, `in-progress`, `done`).
- **Dedicated List Commands**: Quick access to completed, in-progress, or uncompleted tasks.
- **Persistence**: Tasks are automatically saved to a JSON file (`data/tasks.json`).

## Installation

### Prerequisites

- [Go](https://go.dev/doc/install) 1.25 or higher.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/task-tracker.git
   cd task-tracker
   ```

2. Build the application:
   ```bash
   go build -o task-tracker
   ```

3. (Optional) Move the binary to your PATH for global access:
   ```bash
   mv task-tracker /usr/local/bin/
   ```

## Usage

### Add a Task
```bash
task-tracker add "Your task description" --status todo
```

### List All Tasks
```bash
task-tracker list
```

### List Tasks by Status
```bash
task-tracker list --status in-progress
# Or use dedicated commands:
task-tracker list-completed
task-tracker list-inprogress
task-tracker list-uncompleted
```

### Update a Task
```bash
task-tracker update [ID] --desc "Updated description" --status done
```

### Delete a Task
```bash
task-tracker delete [ID]
```

## Project Structure

- `cmd/`: Command-line definitions and CLI logic (using Cobra).
- `internal/task/`: Core business logic, data models, and file persistence.
- `data/`: Directory where `tasks.json` is stored for task persistence.
- `main.go`: Entry point for the application.

## License

Copyright © 2026
