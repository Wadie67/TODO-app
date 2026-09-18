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

	go func() {
		for {
			checkDeadlines()
			time.Sleep(30 * time.Second)
		}
	}()

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
			fmt.Println("Type help!!")
		case "help":
			fmt.Println("Available commands:")
			fmt.Println(" ", "add <task>")
			fmt.Println(" ", "list")
			fmt.Println(" ", "done <task>")
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
		case "done":
			if len(parts) < 2 {
				fmt.Println("Give me a task!")
			} else {
				name := parts[1]
				err := completeTask(name)
				if err != nil {
					fmt.Println("Could not complete task:", err)
				} else {
					fmt.Println("Done with", name)
				}
			}
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
		case "clear":
			err := clearTasks()
			if err != nil {
				fmt.Println("Could not cleartasks", err)
			} else {
				fmt.Println("Cleared!")
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
		today := time.Now()

		parsedTime, err := time.Parse("3:05pm", deadlineinput)
		if err != nil {
			fmt.Println("Not a valid time vro</3")
			continue
		}
		deadline = time.Date(
			today.Year(),
			today.Month(),
			today.Day(),
			parsedTime.Hour(),
			parsedTime.Minute(),
			0,
			0,
			today.Location(),
		)
		fmt.Println("You entered:", deadline.Format("3:04pm"))
		fmt.Println("______________________________________________________________________________________________________")
		break
	}
	return todo{
		Name:      name,
		Priority:  priority,
		Deadline:  deadline,
		Completed: false,
	}

}
