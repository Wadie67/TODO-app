package main

import(
	"time"
	"fmt"
)

type todo struct{
	name string
	priority string	
	deadline time.Time
	completed bool
}

func listtodos() {
	tasks, err := getTasks()
	if err != nil {
		fmt.Println("Could not load tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
		return
	}

	for i, task := range tasks {
		fmt.Println("Task:", i+1)
		fmt.Println("Name:", task.name)
		fmt.Println("Priority:", task.priority)
		fmt.Println("Deadline:", task.deadline.Format("3:04pm"))
		fmt.Println("Completed:", task.completed)
		fmt.Println()
	}
}
