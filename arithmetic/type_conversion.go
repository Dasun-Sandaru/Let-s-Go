package main

import "fmt"

func main() {

	var intNum int = 10
	var floatNum float32 = 3.14

	// sum := intNum + floatNum // (mismatched types int and float32)

	// Type conversion
	sum1 := intNum + int(floatNum)     // 10 + 3 = 13
	sum2 := float32(intNum) + floatNum // 10.0 + 3.14 = 13.14

	fmt.Println(sum1, sum2)
}
