package main

import "fmt"

func main() {
	myslice := []int{1,2,3}

	// Creating a slice from an array
	arr1 := [5]int{1, 3, 5}
	slice1 := arr1[1:3]

	// Creating a slice using make()
	// type, length, capacity
	slice2 := make([]int, 2, 4)
	slice2 = append(slice2, 3, 5)
	slice2 = append(slice2, 7, 9)
	myslice2 := []int{4,5,6}
	slice2 = append(slice2, myslice2...)

	fmt.Printf("Slice 2: %v\n", slice2)
	fmt.Printf("Length of slice 1: %v\n", len(slice1))
	fmt.Printf("Capacity of slice 1: %v\n", cap(slice1))
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
}
