package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

//* Struct
type User struct {
	firstName   string
	lastName    string
	age         int
	createdDate time.Time // Struct within a struct
}

//* Methods - methods are functions within the struct using "receiver arguments"
//? Pointing to User struct
//* Attaching to User struct
// func outputUserDetails(user *User) { // old
// func (Struct) funcName(){... // new
func (user *User) outputUserDetails() {
	//? When you reference dot notation, go assumes you're looking to dereference the fields
	//? The proper notation would be (*user).firstName - (See name values vs age value)
	fmt.Printf("My name is %v %v - age %v\n", user.firstName, user.lastName, (*user).age)
}

func main() {
	//? Requesting pointer to memory address
	var newUser *User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	age := getUserInt("Please enter your age: ")

	//* Created with helper function
	//? Dereference the memory address and saving the value to newUser
	newUser = NewUser(firstName, lastName, age)
	fmt.Println(newUser)
	year, month, day := newUser.createdDate.Date()                                                // y, m, d := myTimestamp.time.Time.Date() - can be destructured into y,m,d
	fmt.Println(newUser.firstName, newUser.lastName, newUser.age, time.Time(newUser.createdDate)) // {Bob Bass 31 {13847557599566623880 2194786001 0x104c9e2a0}}
	fmt.Println(month, day, year)                                                                 // June 21 2021

	//? Access struct keys with dot notation
	newUser.outputUserDetails()

}

//? Gets string
func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	cleanedInput := strings.Replace(userInput, "\n", "", -1)
	return cleanedInput
}

//? Gets int
func getUserInt(promptText string) int {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	cleanedInput := strings.Replace(userInput, "\n", "", -1)
	userAge, _ := strconv.Atoi(cleanedInput)
	return userAge
}

//* Returns pointer to the value created
func NewUser(fName string, lName string, uAge int) *User {
	user := User{
		firstName:   fName,
		lastName:    lName,
		age:         uAge,
		createdDate: time.Now(),
	}
	//? returning pointer value
	return &user
}
