package main

import "fmt"

func main() {
	var (
		firstName, lastName string
		age 			   int
		salary 			   float64
		isConfirmed		   bool
	)
	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T, isConfirmed: %T\n",
		firstName, lastName, age, salary, isConfirmed)

}

