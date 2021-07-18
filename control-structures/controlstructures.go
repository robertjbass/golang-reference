package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserInput(text string) (string, error) {
	fmt.Printf("%v: ", text)
	userInput, err := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1) // -1 for all occurances
	fmt.Println(userInput)
	return userInput, err
}

func getUserAge() (int64, error) {
	userAgeInput, err := getUserInput("Please enter your age")
	if err != nil {
		fmt.Println("Invalid Input")
		return -1, createError("Bad Input")
	}
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)
	if err != nil {
		fmt.Println("Invalid Input")
		return -1, createError("Bad Input")
	}
	return userAge, nil
}

func main() {
	userAge, err := getUserAge()
	if err != nil {
		return
	}
	if userAge >= 30 && userAge < 50 || userAge >= 60 {
		fmt.Println("You're eligable for a senior position.")
	} else if userAge >= 50 {
		fmt.Println("This is the best age.")
	} else if userAge >= 18 {
		fmt.Println("Welcome to the club.")
	} else {
		fmt.Println("Redirecting to Club Penguin.")
	}

	fmt.Println("Please pick an option:")
	fmt.Println("1: Do this\n2: Do this other thing\n3: Do this other other thing")

	userInputOptn, err := getUserInput("Please enter your age")
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}
	if userInputOptn == "1" {
		fmt.Println("Option 1")
	} else if userInputOptn == "2" {
		fmt.Println("Option 2")
	} else if userInputOptn == "3" {
		fmt.Println("Option 3")
	} else {
		fmt.Println("Invalid choice!")
	}

	fmt.Println("Goodbye")

}

func createError(msg string) error {
	return errors.New(msg)
}
