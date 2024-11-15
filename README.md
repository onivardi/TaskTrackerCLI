

Here is the README.md file in Markdown language:


# TaskTrackerCLI
================


A command-line interface for managing tasks.


## Description
---------------


TaskTrackerCLI is a simple command-line tool that allows you to manage your tasks. You can add, delete, and update tasks using a variety of commands.


### Usage
---------


To use TaskTrackerCLI, simply run the `task-cli` command followed by the desired action.


#### Adding a Task

To add a new task, use the `-add` flag followed by the task description.

```bash
task-cli -add "Buy groceries"
```

#### Deleting a Task

To delete a task, use the `-delete` flag followed by the task ID.

```bash
task-cli -delete 1
```

#### Updating a Task



Here is the updated README.md file:


# TaskTrackerCLI
================


A command-line interface for managing tasks.


## Description
---------------


TaskTrackerCLI is a simple command-line tool that allows you to manage your tasks. You can add, delete, update, and list tasks using a variety of commands.


### Usage
---------


To use TaskTrackerCLI, simply run the `task-cli` command followed by the desired action.


#### Adding a Task

To add a new task, use the `-add` flag followed by the task description.

```bash
task-cli -add "Buy groceries"
```

#### Deleting a Task

To delete a task, use the `-delete` flag followed by the task ID.

```bash
task-cli -delete 1
```

#### Updating a Task

To update a task, use the `-update` flag followed by the task ID and the new task description.

```bash
task-cli -update 1 "Buy milk"
```

#### Listing Tasks

To list all tasks, use the `-list` flag.

```bash
task-cli -list
```

#### Listing Tasks by Status

To list tasks by status, use the `-list` flag followed by the status.

```bash
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Commands
------------


Here is a list of available commands:


* `-add`: Add a new task
* `-delete`: Delete a task by ID
* `-update`: Update a task by ID
* `-list`: List all tasks
* `list <status>`: List tasks by status (done, todo, in-progress)

## Installation
---------------


To install TaskTrackerCLI, simply clone the repository and run the `go build` command.

```bash
git clone https://github.com/onivardi/TaskTrackerCLI.git
cd TaskTrackerCLI
go build
```

## Testing
-----------


To run the tests, use the `go test` command.

```bash
go test
```

## License
---------


TaskTrackerCLI is licensed under the MIT License.

## Contributing
------------


Contributions are welcome! If you'd like to contribute to TaskTrackerCLI, please fork the repository and submit a pull request.
