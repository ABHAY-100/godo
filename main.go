package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
    task string;
    completed bool;
}

func main() {
    tasks :=[]Task{}

    for {
        showMenu()
        option := getUserInput("Enter your choice")
    }
}