package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	task      string
	completed bool
}

func showMenu() {
	fmt.Println("\nMenu:")
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

func main() {
	tasks := []Task{}

	for {
		showMenu()

		option := getUserInput("Enter your choice")

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