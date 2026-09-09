package main
import(
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	date := today.Weekday()
	fmt.Println("Hello Wadie!")
	fmt.Println(" ")
	fmt.Println("Today is a", date)
	fmt.Println("What would you like to do today?")
}