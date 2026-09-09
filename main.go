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

		fmt.Println("You want to:", input)

			if input == "quit" {
				break
			}
	}

}