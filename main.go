package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const taskFile = "tasks.json"

func loadTasks() []Task {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(taskFile, data, 0644)
}

func addTask(description string) {
	tasks := loadTasks()
	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func updateTask(id int, newDescription, newStatus string) {
	tasks := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			if newDescription != "" {
				tasks[i].Description = newDescription
			}
			if newStatus != "" {
				tasks[i].Status = newStatus
			}
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Println("Task updated successfully!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func deleteTask(id int) {
	tasks := loadTasks()
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	saveTasks(newTasks)
	fmt.Println("Task deleted successfully!")
}

func listTasks(filter string) {
	tasks := loadTasks()
	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("ID: %d | Description: %s | Status: %s | Created At: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"))
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [command] [arguments]")
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add \"Task Description\"")
			return
		}
		addTask(os.Args[2])

	case "update":
		if len(os.Args) < 5 {
			fmt.Println("Usage: go run main.go update [ID] \"New Description\" [New Status]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID. It must be a number.")
			return
		}
		updateTask(id, os.Args[3], os.Args[4])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete [ID]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID. It must be a number.")
			return
		}
		deleteTask(id)

	case "list":
		listTasks("")

	case "list-done":
		listTasks("done")

	case "list-todo":
		listTasks("todo")

	case "list-inprogress":
		listTasks("in progress")

	default:
		fmt.Println("Unknown command!")
	}
}
