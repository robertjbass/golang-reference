package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//* Printing
	//? fmt is the "format" package
	//? Println() is "Print Line" and reserves a full line ended with a newline
	fmt.Println("hello fmt")

	//* Variable Assignment
	//? var should be used to define a specific type
	var firstString string = "First Var"
	//! `:=` can only be used inside of a function, the full...
	//! `var varName (varType)` syntax must be used ourside of the function
	second := "Second"
	fmt.Println(firstString, second)
	//? var varName varType - must include type without a direct value
	//? var varName = value - will infer type

	//? constants - const instead of var - immutable, must be known when defined
	const pi float32 = 3.14159265358979

	//* Naming Variables for Imports
	//? This variable can not be imported:
	var myLocalVar string = "Because it's lowercase"
	//? This variable can be imported:
	var MyLocalVar string = "Because it's uppercase"
	fmt.Println(myLocalVar, MyLocalVar)

	//* Types

	var typeInt int       // 0
	var typeString string // ""
	var typeFloat float64 // 0.0 (or 0)
	var typeArray [5]int  // [0 0 0 0 0]
	var typeBool bool     // false
	fmt.Println("Empty int, string, float, int array, bool\n", typeInt, "\n", typeString, "\n", typeFloat, "\n", typeArray, "\n", typeBool)

	//* int - Number without decimals (-10 to 151)
	var intNumber float64 = 10 / 3 // 3
	fmt.Println(intNumber)

	//* float64 - Number with decimal places 3.3333333333333335
	//? float64 is heavy on memory
	var floatNumber float64 = float64(10) / 3  // 3.33
	defaultFloat := 1.234567890123456789123456 // Inferred type: float64

	//* float32
	//? less heavy in memory but also less precise
	var smallFloat float32 = 234567890123456789123456 // float32
	smallerFloat := 1.2                               // also defaults to float64
	fmt.Println(floatNumber, defaultFloat, smallFloat, smallerFloat)
	// string - "Hello"

	//* bool
	//? binary yes/no (true or false)
	var firstBool bool = true // true
	fmt.Println(firstBool)

	//* rune
	//? Single unicode character ('❤️' or 'a')
	//? Uses single quotes - larger than a byte, prints an int value
	var firstRune rune = '€' // 8364
	fmt.Println("rune: ", firstRune)
	var runeToString string = string(firstRune)
	fmt.Println("rune to string: ", runeToString)
	defaultRune := 'b' // ends up defaulting to rune, not byte
	fmt.Println("defaultByte ", defaultRune)

	//* byte
	//? Smallest type, single byte ASCII character ('a' or '9')
	var firstByte byte = 'A'  // 65
	var secondByte byte = 'a' // 97
	fmt.Println("byte: ", firstByte, secondByte)
	var firstByteString string = string(firstByte) // A
	fmt.Println("Bytes to Strings: ", firstByteString)

	//* Variables
	//* Naming conventions
	//? Uses mixed case: helloWorld or HelloWorld are both good names

	var greetingText string
	greetingText = "Hello"
	fmt.Println(greetingText)
	// var varName varType
	otherGreeting := "World"
	fmt.Println(otherGreeting)
	// varName := "value" (string type is inferred)

	luckyNumber := 4
	luckiestNumber := luckyNumber
	fmt.Println(luckiestNumber)

	//* Arrays
	//? Array Syntax
	// [length]type{"array", "values"}
	//? Conventions
	// zero indexed

	var arrLen2 [2]string
	arrLen2[0] = "Hello"
	arrLen2[1] = "World"
	fmt.Println(arrLen2[0], arrLen2[1]) // Hello World
	fmt.Println(arrLen2)                // [Hello World]

	arr := [5]int{}              // [1 3 0 0 0]
	odd := [5]int{1, 3, 5, 7, 9} // [0 0 0 0 0]
	twoOfFive := [5]int{1, 3}    // [1 3 5 7 9]
	emptyString := [5]string{}   // ["" "" "" "" ""]
	notDefined := []int{}        // []
	noLen := []int{2, 4, 6}      // [2 4 6]
	fmt.Println(arr, "\n", odd, "\n", twoOfFive, "\n", emptyString, "\n", notDefined, "\n", noLen)

	//* Functions
	//? Functions defined outside of main()

	addResponse := add(4, 5)
	fmt.Println("sum of 4 and 5", addResponse)

	returnArray := returnArray()
	fmt.Println("[2]string{\"Hello\", \"Array\"}")
	fmt.Println(returnArray)

	//? Multiple Returns
	num1, num2 := generateRandomNumbers()
	fmt.Println(num1, num2)

	//? Named Returns
	sum := randomNumberSum(6, 9)
	fmt.Println(sum)

	//* Printing
	//? fmt.Println() - creates a newline at the end
	firstName := "Bob"
	lastName := "Bass"
	fmt.Println(firstName + " " + lastName) // Bob Bass

	firstNumber := 5
	secondNumber := 10
	fmt.Println(firstNumber + secondNumber) // 15 (as an int)
	fmt.Println("15")                       // 15 (as a string)
	//? Mixed types can't be used together, the types need to match
	//? A package can be used to convert "15" to 15
	multiLine := `This is a multiline string
	It can be quite long`
	fmt.Println(multiLine)

	//* Format string print
	//? default formatter or custom, outputs formatted string to standard output
	//? Many of these are taken from C
	// fmt.Print(str)
	// fmt.Printf(format, str)
	// fmt.Println(str)
	fmt.Print("Hello Print") // This is the previous line.Hello Print

	//* Format (Placeholders)
	//? %v - value
	//? Look up the fmt package docs to get all of the placeholder tyoes
	fullName := firstName + " " + lastName
	age := 30
	fmt.Printf("I am %v and I am %v years old.", fullName, age) // I am Bob Bass and I am 30 years old.
	//? Type can be added
	fmt.Printf("I am %v and I am %v (Type: %T) years old.", fullName, age, age) // I am Bob Bass and I am 30 (Type: int) years old.
	fmt.Println("Hello Println")

	//* Sprint
	//? Format string (default formatter or custom) and return formatted string
	//? Returns the value as well
	otherFullName := fmt.Sprintf("%v %v", firstName, lastName) // Bob Bass
	fmt.Println(otherFullName)

	//? Sprint types
	// fmt.Sprint(str)
	// fmt.Sprintf(format, str)
	// fmt.Sprintln(str)

	//* Fprint
	//? Format string (default formatter or custom) and write to specified writer
	//? Sprint types
	// fmt.Fprint(writer, str)
	// fmt.Fprintf(writer, format, string)
	// fmt.Fprintln(writer, str)

	//* ==================
	//* POINTERS
	//* ==================

	//* Pointers
	thisName := "Bob"     // computer memory: 0xc0000018050
	thisAge := 30         // computer memory: 0xc0110020011
	userAge := &age       // 0xc0110020011
	userName := &thisName // 0xc0000018050
	fmt.Println(thisName, thisAge, userAge, userName)

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

//* ====================
//*	External Functions
//* ====================

//* Function structure
// func name(varName varType) returnType {
// 	return varName
// }

func add(a int, b int) int {
	sum := a + b
	return sum
}

func returnArray() [2]string {
	return [2]string{"Hello", "Array"}
}

func generateRandomNumbers() (int, int) {
	fmt.Println("multiple returns")
	randomNumber1 := rand.Intn(10)
	randomNumber2 := rand.Intn(10)
	return randomNumber1, randomNumber2
}

func randomNumberSum(firstNumMax int, secondNumMax int) (sum int) {
	fmt.Println("return named value")
	randomNumber1 := rand.Intn(firstNumMax)
	randomNumber2 := rand.Intn(secondNumMax)
	sum = randomNumber1 + randomNumber2
	// return sum is implied
	return
}

//* ==================
//* POINTERS
//* ==================

//* Better performance
func double(number *int) int {
	result := *number * 2
	*number = 100 // this will update the value of the number to 100
	return result
}

// todo - Structs and Pointers are in external files
