package main

import "fmt"

func main() {
	var name = "Sam"
	var highLevelLanguagesKnown, lowLevelLanguagesKnown = 12, 3
	var junkInteger int
	colorName := "Red"
	fmt.Println("Welcome to the variables Go Program")
	fmt.Println("Let's explore the variables one by one")
	fmt.Println("My name is ", name)
	fmt.Println("The total numbers of programming languages I know is", highLevelLanguagesKnown+lowLevelLanguagesKnown)
	fmt.Println("The default value of undefined int variable is", junkInteger)
	fmt.Println("We can declare variables using the := symbol, the value of colorName is", colorName)

}
