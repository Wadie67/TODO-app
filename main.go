package main
import(
	"fmt"
	"time"
	"os"
	"bufio"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	today := time.Now()
	date := today.Weekday()

	fmt.Println("Hello Wadie!")
	fmt.Println(" ")
	fmt.Println("Today is a", date)
	fmt.Println("What would you like to do today?")

	for scanner.Scan() {
		input := scanner.Text()

		if input == "-help" {
			fmt.Println("Available commands:")
			fmt.Println(" ", "add <task>")
			fmt.Println(" ", "list")
			fmt.Println(" ", "done <number>")
			fmt.Println(" ", "delete <number>")
			fmt.Println(" ", "clear")
			fmt.Println(" ", "help")
			fmt.Println(" ", "quit")
		} 
		if input == "quit" {
				break
			} else {
			fmt.Println("You want to:", input)
		}
	}

}