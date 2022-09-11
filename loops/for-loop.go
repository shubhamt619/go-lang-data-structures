package main

import "fmt"

func main() {
	fmt.Println("For Loop Example")
	messagesSent := 0

	for messagesSent < 10 {
		fmt.Printf("Message number %v sent\n", messagesSent)
		messagesSent++
	}
}
