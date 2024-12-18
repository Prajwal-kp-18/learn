# Task CLI

Task CLI is a simple command-line interface application written in Go that allows users to manage tasks. You can add, list, complete, and delete tasks, with all data stored in a JSON file for persistence.

## Features
- Add new tasks with a title.
- List all tasks with their status (completed or not).
- Mark tasks as completed.
- Delete tasks.

## Prerequisites
- [Go](https://golang.org/) installed on your system (version 1.16 or later).

## Installation
1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_folder>
   ```

2. Build the application:
   ```bash
   go build -o learn
   ```

3. Move the binary to a directory in your PATH for easy access (optional):
   ```bash
   mv learn /usr/local/bin/
   ```

## Usage
### General Command Format
```bash
learn <command> [arguments]
```

### Commands
1. **Add a Task**
   Add a new task by specifying its title:
   ```bash
   learn add "Task title"
   ```
   Example:
   ```bash
   learn add "Buy groceries"
   ```

2. **List All Tasks**
   Display all tasks along with their statuses (completed or not):
   ```bash
   learn list
   ```

3. **Mark a Task as Completed**
   Mark a specific task as completed by its ID:
   ```bash
   learn complete <task_id>
   ```
   Example:
   ```bash
   learn complete 1
   ```

4. **Delete a Task**
   Delete a specific task by its ID:
   ```bash
   learn delete <task_id>
   ```
   Example:
   ```bash
   learn delete 1
   ```

### Example Workflow
1. Add tasks:
   ```bash
   learn add "Read a book"
   learn add "Write Go code"
   ```

2. List tasks:
   ```bash
   learn list
   ```
   Output:
   ```
   1. Read a book [❌]
   2. Write Go code [❌]
   ```

3. Mark the first task as completed:
   ```bash
   learn complete 1
   ```

4. List tasks again:
   ```bash
   learn list
   ```
   Output:
   ```
   1. Read a book [✅]
   2. Write Go code [❌]
   ```

5. Delete a task:
   ```bash
   learn delete 2
   ```

## File Structure
- **tasks.json**: The application stores task data in a `tasks.json` file in the current directory. If the file does not exist, it will be created automatically.

## Error Handling
- If an invalid task ID is provided for `complete` or `delete` commands, the application will notify you that the task was not found.
- If no tasks are present, the `list` command will notify you accordingly.

## Contribution
Feel free to fork the repository, submit issues, and make pull requests. Contributions are always welcome!

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Author
Prajwal-kp-18 

