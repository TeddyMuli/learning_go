package main

import "fmt"

func main() {
	// Specifying length
	var arr1 = [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	// Infering the length of the array
	var arr3 = [...]string{"ted", "muli"}
	arr4 := [...]string{"goat", "sheep", "cow"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
