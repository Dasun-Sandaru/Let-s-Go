package main

import "fmt"

func main() {
	a := 10
	b := 3

	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	remainder := a % b

	fmt.Println(sum, sub, mul, div, remainder)

	c := 5
	c = c + 2
	c += 2 // c = c + 2 = 7
	c -= 2 // c = c - 2 = 5
	c *= 2 // c = c * 2 = 10
	c /= 2 // c = c / 2 = 5
	c %= 2 // c = c % 2 = 1

	fmt.Println(c)

	d := 10
	d++ // d = d + 1 = 11
	d-- // d = d - 1 = 10
	fmt.Println(d)
}
