package main

// increment increments the value of the variable
func increment(value *int) {
	*value++
}

// setToZero sets the value of the variable to zero
func setToZero(value *int) {
	*value = 0
}

func main() {
	num := 10
	increment(&num)
	println("After incrementing and setting to zero: ", num)

	setToZero(&num)
	println("After setting to zero: ", num)
}
