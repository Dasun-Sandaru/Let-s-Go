package main

import "fmt"

func main() {

	fmt.Println("Dasun Sandaru")
	fmt.Println("Dasun Sandaru")
	fmt.Println("Dasun Sandaru")

	fmt.Print("Dasun")
	fmt.Print(" Sandaru")
	fmt.Println()

	fmt.Println("-----------")

	fmt.Println(10 + 30)
	fmt.Println(10 - 20)

	fmt.Println("-----------")

	var firstNumber = 10  // var firstNumber int = 10
	var secondNumber = 20 // var secondNumber int = 20

	fmt.Println(firstNumber + secondNumber)
	fmt.Println(firstNumber + secondNumber)
	fmt.Println(firstNumber - secondNumber)

	fmt.Println("-----------")

	var firstName = "Dasun" // var firstName string = "Dasun"
	print(firstName)

	fmt.Println("-----------")

	var userScore = 25.5      // var userScore float32 = 25.5 or var userScore float64 = 25.5
	var userBalance = -100.00 // var userBalance float32 = -100.00 or var userBalance float64 = -100.00

	fmt.Println(userScore, userBalance) // 25.5 -100.00

	fmt.Println("-----------")
	fmt.Printf("User Score: %T\n", userScore)

	fmt.Println("-----------")
	var isUserFound = true // var isTrue bool = true

	fmt.Println(isUserFound)

	fmt.Println("-----------")
	var userAge, userWeight = 25, 75.5 // var userAge int = 25, userWeight float32 = 75.5
	fmt.Println(userAge, userWeight)

	fmt.Println("-----------")

	newUserName := "Sandaru" // var newUserName string = "Sandaru"
	fmt.Println(newUserName)

}
