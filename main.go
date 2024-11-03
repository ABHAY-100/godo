package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	text string
	completed bool
}

func showMenu() {
	time.Sleep(1 * time.Second)
	fmt.Println("\n\nMenu:")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add Task")
	fmt.Println("3. Mark Task as Completed")
	fmt.Println("4. Save Tasks to File")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("**No tasks available**")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := " "
		if task.completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.text)
	}
}

func addTask(tasks *[]Task) {
	taskText := getUserInput("Enter task description: ")
	*tasks = append(*tasks, Task{text: taskText})
	fmt.Println("Task added.")
}

func markTaskCompleted(tasks *[]Task) {
	showTasks(*tasks)
	taskIndexStr := getUserInput("\nEnter task number to mark as completed: ")
	taskIndex, err := strconv.Atoi(taskIndexStr)
	if err != nil || taskIndex < 1 || taskIndex > len(*tasks) {
	 fmt.Println("Invalid task number. Please try again.")
	 return
	}
	(*tasks)[taskIndex-1].completed = true
	fmt.Println("Task marked as completed.")
}

func saveTasksToFile(tasks []Task) {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	for _, task := range tasks {
		status := " "
		if task.completed {
			status = "x"
		}
		file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.text))
	}
	fmt.Println("Tasks saved to file 'tasks.txt'.")
}

func loadTasksFromFile() ([]Task, error) {
	file, err := os.Open("tasks.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		status := line[1]
		taskText := line[4:]
		completed := status == 'x'
		tasks = append(tasks, Task{text: taskText, completed: completed})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func main() {
	tasks, err := loadTasksFromFile()
	if err != nil {
		fmt.Println("Error loading tasks from file:", err)
		tasks = []Task{}
	}

	for {
		showMenu()

		option := getUserInput("\nEnter your choice: ")

		switch option {
            case "1": 
                showTasks(tasks)
            case "2": 
                addTask(&tasks)
            case "3": 
                markTaskCompleted(&tasks)
            case "4": 
                saveTasksToFile(tasks)
            case "5": 
                fmt.Println("Exiting the ToDo application.")
                return
            default:
                fmt.Println("Invalid choice. Please try again")
		}
	}
}
