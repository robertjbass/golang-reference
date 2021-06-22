package main

import "fmt"

func main() {
	age := 30
	fmt.Println("age: ", age)

	//* & tells go to get the memory address of age and store it in myAge
	//? type: var myAge *int - asterisk indicates pointer
	myAge := &age // 0x140000b2000

	//? When fully written out
	// var myAge2 *int
	// myAge2 = &age // 0x140000b2000

	fmt.Println("age memory address: ", myAge)        // 0x140000b2000
	fmt.Println("age memory address value: ", *myAge) // 30 // this is called 'dereferencing'

	//! myAge = 30 // this won't work, myAge has already been defined as a pointer

	*myAge = 31

	fmt.Println(*myAge) // 31

	//* the original address is now modified
	fmt.Println(age) // 31

	// double() takes a reference as a value
	doubledAge := double(&age)
	fmt.Println("age", age)               // 31
	fmt.Println("doubledAge", doubledAge) // 62
}

//* Better performance
func double(number *int) int {
	result := *number * 2
	*number = 100 // this will update the value of the number to 100
	return result
}
