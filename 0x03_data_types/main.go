package main

import "fmt"

func main() {
	// Specifying length
	var arr0 = [4]int{} // Empty array
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6}

	// Infering the length of the array
	var arr3 = [...]string{"ted", "muli"}
	arr4 := [...]string{"goat", "sheep", "cow"}

	// Change value
	arr2[4] = 9

	// Initialize specific elements
	var arr5 = [4]int{1:10, 3:7}

	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

	// Find the lenght of an array
	fmt.Println(len(arr4))
}
