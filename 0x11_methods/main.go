package main

import "fmt"

// Struct type receiver
type person struct {
	name string
	age int
}

func (p person) display() {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
}

// Non struct type receiver
type number int

func (n number) square() number {
	return n * n
}

// Method with pointer receivers
func (person *person) changeName(newName string) {
	person.name = newName
}

func main() {
	fmt.Println("1. Struct type receiver")
	a := person{name: "Teddy", age: 21}
	a.display()

	fmt.Println("\n2. Non struct type receiver")
	b := number(4)
	c := b.square()
	fmt.Printf("The square of %d is %d", b,  c)

	fmt.Println("\n\n3. Method with pointer receiver")
	a.changeName("Muli")
	a.display()
}
