package main

import (
	"fmt"
	"strconv"
)

const name string = "sam"

func main() {
	fmt.Println("Name is " + name)

	const multiply_factor = 3

	fmt.Println("Multiply factor is " + strconv.Itoa(multiply_factor))

	const multiplication = multiply_factor * 2

	fmt.Println("Multiplication is " + strconv.Itoa(multiplication))

}
/*
Let us consider that we are moving ahead with shopify as our core framework.


*/