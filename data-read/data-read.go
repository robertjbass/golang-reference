package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/716green/academind/data-read/uuid"
)

var reader = bufio.NewReader(os.Stdin)

type Data struct {
	id          uuid.UUID
	name        string
	description string
	price       float64
}

func generateUid() uuid.UUID {
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return *uid
}

func main() {

	inputName := getUserData("What is the name of the item? ")
	inputDesc := getUserData("What is the description for this item? ")
	inputPrice := getUserData("What is the price of this item? ")

	//* Creating direct struct
	newData := Data{
		id:          generateUid(),
		name:        "Bad Item",
		description: "Created directly in main function",
		price:       0.99,
	}
	fmt.Println(newData)

	//* userData is the value of the memory address returned by createData
	inputPriceFloat, _ := strconv.ParseFloat(inputPrice, 64)
	userData := *createData(inputName, inputDesc, inputPriceFloat)
	fmt.Println(userData)
	Data.ReadData(newData)
	Data.ReadData(userData)

	storeData()
}

func createData(name string, description string, price float64) *Data {
	data := Data{
		id:          generateUid(),
		name:        name,
		description: description,
		price:       float64(price),
	}
	return &data
}

func (data Data) ReadData() {
	fmt.Println(data.id)
	fmt.Println(data.name)
	fmt.Println(data.description)
	fmt.Printf("%0.2f\n", data.price)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	cleanedInput := strings.Replace(userInput, "\n", "", -1)
	return cleanedInput
}

func storeData() {
	fmt.Println("TODO - Create a function that stores the user data as a local file")
}
