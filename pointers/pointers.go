package main

import "fmt"

func main() {

	// Declaring a variable
	var num int = 10
	fmt.Println("Value of num is: ", num)

	// Declaring a pointer
	var ptr *int = &num
	fmt.Println("Value of ptr is: ", ptr)

	// Pointer dereferencing
	fmt.Println("Value of *ptr is: ", *ptr)

	// Changing the value of num using pointer
	*ptr = 20
	fmt.Println("Value of num is: ", num)

	// Pointer Address
	fmt.Println(&ptr)

}
