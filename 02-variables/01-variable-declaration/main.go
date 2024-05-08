package main

import "fmt"

func main() {
	// Declare a variable
	var myStr string = "Lucifer"
	var myInt int = 3600
	var myFloat float64 = 3.14
	fmt.Println(myStr, myInt, myFloat)

	// Declare multiple variables
	var name, age, salary = "Lucifer", 3600, 3.14
	fmt.Println(name, age, salary)

	fullname := "Lucifer"
	favnum := 3600
	remsalary := 3.14
	fmt.Println(fullname, favnum, remsalary)
}