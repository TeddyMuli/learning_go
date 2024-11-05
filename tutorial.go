package main

import "fmt"

func main() {
	fmt.Printf("Hello world!\n")

	var name string
	var age uint
	var option1 string
	var option2 int
	correct_count := 0
	question_count := 0

	fmt.Printf("Enter your name > ")
	fmt.Scan(&name)
	fmt.Println("Goodmorning: ", name)

	fmt.Printf("Enter your age > ")
	fmt.Scan(&age)
	if (age < 18) {
		fmt.Println("You are underage!")
		return
	} else {
		fmt.Println("Welcome!")
	}
	fmt.Println("Continue")

	fmt.Printf("Which is better Go or Python? > ")
	fmt.Scan(&option1)
	question_count++

	if option1 == "Go" || option1 == "go" {
		fmt.Println("Correct")
		correct_count++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many engineers developed Go? > ")
	fmt.Scan(&option2)
	question_count++

	if option2 == 3 {
		fmt.Println("Correct")
		correct_count++
	} else {
		fmt.Println("Incorrect!")
	}

	percentage := (float64(correct_count) / float64(question_count)) * 100

	fmt.Printf("Yo have gotten %v out of %v questions. ", correct_count, question_count)
	fmt.Printf("%v%%", percentage)
}
