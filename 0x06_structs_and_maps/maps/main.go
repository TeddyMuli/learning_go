package main

import "fmt"

func main() {
	var a = map[string]int{"One": 1, "Two": 2}
	// Empty map, 
	var b = make(map[string]string)

	for key, value := range a {
		fmt.Println(key, value)
	}
	fmt.Println(b == nil)
	// Append
	b["Teddy"] = "Muli"

	// Check for something in a map
	_, ok := b["Tedy"]
	fmt.Println(b)

	fmt.Println(ok)
}
