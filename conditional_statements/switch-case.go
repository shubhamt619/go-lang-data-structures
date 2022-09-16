package main

import "fmt"

func main() {
	favouriteAnimal := 2
	fmt.Println("Favourite Animal is at position ", favouriteAnimal)
	switch favouriteAnimal {
	case 1:
		fmt.Println("I like snakes")
	case 2:
		fmt.Println("I like turtles")
	case 3:
		fmt.Println("I like cats")
	}
}
