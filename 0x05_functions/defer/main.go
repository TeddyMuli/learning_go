package main

import "fmt"

func main() {
	defer fmt.Println("world!")
	defer fmt.Println("Ted")
	fmt.Println("Hello")
}
