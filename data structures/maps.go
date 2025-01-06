package main

import "fmt"

func main() {

	// declare a map
	var colors map[string]string = map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	fmt.Println(colors)

	// add elements to a map
	colors["white"] = "#FFFFFF"
	fmt.Println(colors)

	// delete elements from a map
	delete(colors, "green")

	// length of a map
	fmt.Println(len(colors))
}
