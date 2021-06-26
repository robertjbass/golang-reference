package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Struct
type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func checkForFolder(folderName string) {
	if _, err := os.Stat("/" + folderName); os.IsNotExist(err) {
		fmt.Println("'products' folder doesn't exist, creating...")
		os.Mkdir("products", 0755)
	}

	if _, err := os.Stat("/" + folderName); !os.IsNotExist(err) {
		fmt.Println("'products' folder found")
	}
}

// Append method to Product strut
func (prod *Product) store() {
	currentDir, _ := os.ReadDir(".")
	fmt.Println(currentDir)

	checkForFolder("products")

	file, _ := os.Create("products/" + prod.id + ".json")
	prodFolder, _ := os.ReadDir("products")
	fmt.Printf("%v Products\n", len(prodFolder))

	// Created
	content := fmt.Sprintf("{\n\"id\": \"%v\",\n\"title\": \"%v\",\n\"description\": \"%v\",\n\"price_usd\": \"%.2f\"\n}",
		prod.id,
		prod.title,
		prod.description,
		prod.price,
	)

	file.WriteString(content)
	file.Close()
}

// This function connects to the Product struct. It can be called with product.functionName()
func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: $%.2f\n", prod.price)
}

func NewProduct(t string, d string, p float64) *Product {
	u := uuid.New().String()
	return &Product{u, t, d, p}
}

func main() {
	savedProduct := getProduct()

	savedProduct.printData()
	savedProduct.store()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data.")
	fmt.Println("------------------------------")
	reader := bufio.NewReader(os.Stdin)

	titleInput := readUserInput(reader, "Enter Title: ")
	descriptionInput := readUserInput(reader, "Enter Description: ")
	priceInput := readUserInput(reader, "Enter Price: ")

	floatPrice, _ := strconv.ParseFloat(priceInput, 64)

	productInput := NewProduct(titleInput, descriptionInput, floatPrice)

	return productInput
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}
