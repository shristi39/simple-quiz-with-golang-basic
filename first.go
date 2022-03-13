package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Println("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age:")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("you can play!")
	} else {
		fmt.Println("You cannot play!")
		return

	}
	fmt.Println("What is better, the blue or green? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, answer2)

	fmt.Println("Select from 1 to 10")
	var cores int
	fmt.Scan(&cores)

	fmt.Println("What is better, school or college? ")
	var solution string
	var solution2 string
	fmt.Scan(&solution, solution2)

	fmt.Println("What is your favourite, web development or graphics? ")
	var view string
	var view2 string
	fmt.Scan(&view, view2)

	fmt.Println("Thank you for playing this game!")
}
