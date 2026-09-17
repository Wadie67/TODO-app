package main

import(
	"time"
	"fmt"
)

type todo struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Priority  string
	Deadline  time.Time
	Completed bool
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
		fmt.Println("Name:", task.Name)
		fmt.Println("Priority:", task.Priority)
		fmt.Println("Deadline:", task.Deadline.Format("3:04pm"))
		fmt.Println("Completed:", task.Completed)
		fmt.Println()
	}
}
