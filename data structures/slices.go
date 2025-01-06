package main

import "fmt"

func main() {

	// declare and intialize a slice

	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{100,200,300,400,500}
	fmt.Println(numbers1)
	fmt.Println(numbers2)

	// append elements to a slice
	numbers1 = append(numbers1, 10, 11, 12)
	fmt.Println(numbers1)
	fmt.Println(len(numbers1))

	// remove elements from a slice
	numbers2 = append(numbers2[:2], numbers2[3:]...)
	fmt.Println(numbers2)

}
