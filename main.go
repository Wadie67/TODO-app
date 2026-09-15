package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	today := time.Now()
	date := today.Weekday()
	weather := getWeather()
	err := initDB()

	if err != nil {
		fmt.Println("Database error:", err)
		return
	}
	defer db.Close()

	fmt.Println("Hello Wadie!")
	fmt.Println(" ")
	fmt.Println("Today is a", weather, date)
	fmt.Println("What would you like to do today?")
	fmt.Print("> ")
loop:
	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		switch parts[0] {
		default:
			fmt.Println("You want to:", parts[1])
		case "help":
			fmt.Println("Available commands:")
			fmt.Println(" ", "add <task>")
			fmt.Println(" ", "list")
			fmt.Println(" ", "done <number>")
			fmt.Println(" ", "delete <task>")
			fmt.Println(" ", "clear")
			fmt.Println(" ", "help")
			fmt.Println(" ", "quit")
		case "add":
			if len(parts) < 2 {
				fmt.Println("Give me a task!")
			} else {
				name := parts[1]
				newtodo := createtodo(scanner, name)
				err := addTask(newtodo)
				if err != nil {
					fmt.Println("Could not save task:", err)
				} else {
				fmt.Println("Task saved!")
				}
			}
		case "list":
			listtodos()
		case "delete":
			if len(parts) < 2 {
				fmt.Println("Give me a task!")
			} else {
				name := parts[1]
				err := deleteTask(name)
				if err != nil {
					fmt.Println("Could not delete task:", err)
				} else {
					fmt.Println("Deleted", name)
				}
			}
		case "quit":
			fmt.Println("Goodbye!")
			break loop
		}
		fmt.Print("> ")
	}
}
func createtodo(scanner *bufio.Scanner, name string) todo {

	fmt.Println("What is the Priority of this Task?")
	fmt.Print("> ")
	scanner.Scan()
	priority := scanner.Text()
	fmt.Println("You entered:", priority)

	fmt.Println("When would you like this Task done?")
	var deadline time.Time
	for {
		fmt.Print("> ")
		scanner.Scan()
		deadlineinput := scanner.Text()
		deadline, err := time.Parse("3:05pm", deadlineinput)
		if err != nil {
			fmt.Println("Not a valid time vro</3")
			continue
		} else {
			fmt.Println("You entered:", deadline.Format("3:05pm"))
			fmt.Println("______________________________________________________________________________________________________")
			break
		}
	}
		return todo{
		name:      name,
		priority:  priority,
		deadline:  deadline,
		completed: false,
	}
}
