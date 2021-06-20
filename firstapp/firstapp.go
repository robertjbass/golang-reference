package main

import (
	"fmt"

	"github.com/716green/academind/firstapp/greeting"
)

func main() {
	fmt.Println("Hello World")

	// greetingText := "Hello World"
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greeting.GreetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)
}
