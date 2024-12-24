package main
import (
	"fmt"
	"time"
)

var day string

func main() {
	fmt.Printf("Today is %s", day)
}

func init() {
	day = time.Now().Weekday().String()
}

func init() {
	fmt.Println("First init")
}

func init() {
	fmt.Println("Second init")
}

func init() {
	fmt.Println("Third init")
}

func init() {
	fmt.Println("Fourth init")
}
