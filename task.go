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


var todos = make(map[string]todo)

func listtodos() {
    for _, task := range todos {
        fmt.Println("Task:", task.name)
        fmt.Println("Priority:", task.priority)
        fmt.Println("Deadline:", task.deadline.Format("3:04pm"))
        fmt.Println("Completed:", task.completed)
        fmt.Println()
    }
}
