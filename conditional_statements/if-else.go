package main

import "fmt"

func main() {
	if 10%2 == 0 {
		fmt.Println("10 is divisible by 2")
	} else {
		fmt.Println("10 is not divisible by 2")
	}

	if 10 < 100 {
		fmt.Println("10 is less than 100")
	} else if 10 < 9 {
		fmt.Println("10 is less than 9")
	} else {
		fmt.Println("10 is greater than 100")
	}
}
