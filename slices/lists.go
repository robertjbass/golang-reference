package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
	id    string
	title string
	price float64
}

func create(t string, p float64) Product {
	newProduct := Product{
		id:    uuid.New().String(),
		title: t,
		price: p,
	}
	return newProduct
}

func (product Product) log() {
	fmt.Println(product)
}

func printDivider() {
	fmt.Println("-----------------")
}

func printSection(title string) {
	fmt.Println("")
	printDivider()
	fmt.Println(title)
	printDivider()
}

func main() {
	myHobbies := [3]string{"Programming", "Working", "Learning"}
	fmt.Println(myHobbies)                // [Programming Working Learning]
	fmt.Println("length", len(myHobbies)) // length 3
	fmt.Println(myHobbies[0])             // Programming
	fmt.Println(myHobbies[1:])            // [Working Learning]
	frontHobbies := myHobbies[:2]
	fmt.Println(frontHobbies) // [Programming Working]
	endHobbies := frontHobbies[1:3]
	fmt.Println(endHobbies) // [Working Learning]
	myGoals := []string{}
	myGoals = append(myGoals, "Learn Golang", "Become a better programmer")
	fmt.Println(myGoals) // [Learn Golang Become a better programmer]
	myGoals[1] = "Learn more about C"
	myGoals = append(myGoals, "Become a better programmer")
	fmt.Println(myGoals) // [Learn Golang Learn more about C Become a better programmer]

	myNewProduct := create("Go Course", 9.99)
	myNewProduct.log()
	allProducts := []Product{myNewProduct}

	allProducts = append(allProducts, create("Typescript Course", 12.99), create("NuxtJS", 11.99))
	allProducts[1].log()
	allProducts[2].log()

	printSection("Assignment Result")
	fmt.Println(allProducts)

	printSection("Maps")
	prices := []float64{9.99, 12.99}
	fmt.Println(prices[0:1]) // 9.99
	prices[1] = 11.99

	prices = append(prices, 13.99, 14.99, 15.99, 16.99)
	fmt.Println(prices)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.12, 201.12, 301.12}
	// Just like JS where we may do:
	// arr2 = [...arr1, ...arr2]
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// 1) Create a new array that contains three hobbies you have
// Output (print) that array in the command line.
// 2) Also output more data about the array:
//    a) The first element (standalone)
//    b) The second and third element combined as a new list
// 3) Create a slice based on the first element that conains the fist and second elements
// 4) Re-slice the slice from (3) and change it to contain the second and last element of the original array.
// 5) Create a "dynamic array" that contains your goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing...
// 7) Bonus:
//    a) Create a "Product" struct with title, id, price, and create a dynamic list of products (at least 2 products).
//    b) Then add a third product to the existing list of products.
