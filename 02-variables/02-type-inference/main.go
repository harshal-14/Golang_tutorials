package main

import "fmt"

func main() {
	// Type inference

	var name = "Ranjit Patel"
	fmt.Printf("Variable 'name' is of type %T\n", name)

	// Multiple variable declaration with inferred types
	var firstName, lastName, age, salary = "Ranjit", "Patel", 25, 3600
	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
	firstName, lastName, age, salary)
}