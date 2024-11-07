package main

import "fmt"

type Person struct {
	name string;
	age int;
	weight float32;
	height float32;
}

func printPerson(person Person) {
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)
	fmt.Println("Weight: ", person.weight)
	fmt.Println("Height: ", person.height)
}

func main() {
	var person1 Person
	var person2 Person

	person1.name = "Teddy"
	person1.age = 21
	person1.weight = 54
	person1.height = 165

	person2.name = "Jane"
	person2.age = 70
	person2.weight = 65
	person2.height = 155

	printPerson(person1)
	printPerson(person2)
}
