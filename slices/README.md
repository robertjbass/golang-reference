# Array - List - Slice

## Array Syntax

```go
// [howMany]type{value1, value2}
prices := [2]float64{1.99, 2.24}

// Set Empty
var productNames [4]string // ["" "" "" ""]

```

## Slices
```go
package main

import (
	"fmt"
)

func forEach(arr [12]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {

	months := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	summer := months[6:8]

	fmt.Println(summer)

	forEach(months)

	fmt.Println("len", len(months), "cap", cap(months)) // 12, 12
	fmt.Println("len", len(summer), "cap", cap(summer)) // 2, 6

}

```

### More Slice Examples

```go
	arrPrice := [2]float64{12.34, 56.78}
	slicePrice := []float64{34.56, 78.90}
	prices := []float64{12.34, 56.78, 90.12, 34.56, 78.90}
	sliceOfPrice := prices[2:5]
	fmt.Println(arrPrice, slicePrice, prices, sliceOfPrice)
	updatedPrices := append(prices, 9.99)
	fmt.Println(updatedPrices)
	fmt.Println(prices)
```

# Assignment

1) Create a new array that contains three hobbies you have
Output (print) that array in the command line.
2) Also output more data about the array:
   a) The first element (standalone)
   b) The second and third element combined as a new list
3) Create a slice based on the first element that conains the fist and second elements
4) Re-slice the slice from (3) and change it to contain the second and last element of the original array.
5) Create a "dynamic array" that contains your goals (at least 2 goals)
6) Set the second goal to a different one AND then add a third goal to that existing...
7) Bonus:
   a) Create a "Product" struct with title, id, price, and create a dynamic list of products (at least 2 products).
   b) Then add a third product to the existing list of products.

```go
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
	printDivider()

	fmt.Println(allProducts)
}
```