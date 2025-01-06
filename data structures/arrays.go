package main

import "fmt"

func main() {

	var numbers1 [5]int = [5]int{1, 2, 3, 4, 5} // ok
	//var numbers2 = [5]int{1, 2, 3, 4, 5}      // ok

	fmt.Println(numbers1)
	fmt.Println(numbers1[0])

	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	fmt.Println("Length of colors array is", len(colors))

	// Change the value of an element
	colors[0] = "Yellow"
	fmt.Println(colors)

	// Arrary length is unknown
	var numbers3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers3)

	// numbers3[5] = 6 // Error: invalid array index 5 (out of bounds for 5-element array)

	// 2D array
	var matrix = [2][3]int{{1, 2, 3}, {4, 5, 6}} // 2 rows, 3 columns
	fmt.Println(matrix) // [[1 2 3] [4 5 6]]
}
