package main

import "fmt"

func main() {
	defer fmt.Println("world!")
	defer fmt.Println("Ted")
	fmt.Println("Hello")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
