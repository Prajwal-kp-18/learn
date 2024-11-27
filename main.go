package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Completed bool   `json:"completed"`
}

var taskFile = "tasks.json"

// Load tasks from the file
func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := ioutil.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil 
		}
		return nil, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// Save tasks to the file
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(taskFile, data, 0644)
}

// Add a new task
func addTask(title string) {
	tasks, _ := loadTasks()
	id := len(tasks) + 1
	task := Task{ID: id, Title: title, Completed: false}
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Println("Task added!")
}

// List all tasks
func listTasks() {
	tasks, _ := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
	}
}

// Mark a task as completed
func completeTask(id int) {
	tasks, _ := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			saveTasks(tasks)
			fmt.Println("Task marked as completed!")
			return
		}
	}
	fmt.Println("Task not found!")
}

// Delete a task
func deleteTask(id int) {
	tasks, _ := loadTasks()
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	saveTasks(newTasks)
	fmt.Println("Task deleted!")
}

// Main function to handle commands
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: learn <command> [arguments]")
		fmt.Println("Commands: add, list, complete, delete")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task add <task_title>")
			return
		}
		title := os.Args[2]
		addTask(title)
	case "list":
		listTasks()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task complete <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		completeTask(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task delete <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		deleteTask(id)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Commands: add, list, complete, delete")
	}
}
