# Task Tracker

Task Tracker is a command-line interface (CLI) application for managing tasks. It allows you to add, update, delete, mark tasks as in progress or done, and list tasks based on their status.

## Features
- **Add**: Add a new task with a description and a default status of "todo",
- **List**: View tasks based on their status (todo, in progress, done, or all).
- **Update**: Update the description of an existing task based on their ID.
- **Mark**: Change the status of a task into "in progress" or "done" based on their ID.
- **Delete**: Remove an existing task based on their ID.

## Commands

### Adding a New Task

To add a new task, use the following command:

```sh
go run main.go add "Buy groceries"
```
### Update a Task

To update a task, use the following command:

```sh
go run main.go update 1 "Buy groceries and cook dinner"
```

### Delete a Task

To remove a task, use the following command:

```sh
go run main.go remove 1
```

### Mark as Inprogress
there is 3 status : [todo , in-progress,done]

To change status of a task, use the following command:

```sh
go run main.go mark-in-progress 1
```
```sh
go run main.go mark-done 1
```

### Get list of Tasks

to get list of tasks:

```sh
go run main.go list
```

to get list with status filter:
```sh
go run main.go list in-progress
```
