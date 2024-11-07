package main

import "fmt"

func main() {
	arr1 := [3]int{1, 3,4}

	// for loop for arrays and slices using range
	for _, value := range(arr1) {
		fmt.Printf("%v\n", value)
	}
}
