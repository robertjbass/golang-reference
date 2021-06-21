# Golang

## Syntax, Formats, Conventions, Standards, and Examples

### Each subfolder contains its own go.mod file and are treated as separate projects

> This repo is to document, explain, and reference the features of the Go programming language

  > The purpose of this repo is to create a series of references that I can use in the future when working with the Go programming language. As a (mostly) JavaScript developer, these references are important to compare and contrast language features and syntax

* There is a `.go` file in `./notes/syntax.go` that contains all of the examples from this `README.md` file.
* Special thanks to Maximillian Schwartzmüller of Academind for creating an excellent course on Golang 👏

📁 `firstapp` is a self-contained module within this repo. It is used to demonstrate go.mod and package structures

📁 `notes` contains code snippets contained in this `README.md` file that can be run with the Go compiler

📁 `assignments` contains exercises that I'm using to test how well I'm retaining the concepts without referencing the notes

## General Go Notes
* Go uses semicolons, they are omitted in development and created behind the scenes
* The modules should match the host URL, for example
  * File Structure: 📁 github.com/716green/myGoApp
  * Repo: https://github.com/716green/myGoApp
* Go looks for a `main` package and `main` func when using `go run .`


## Go File Structure

```go

package main

import "fmt"

func main() {
  fmt.Println("Hello World")
}
```

## Printing

- fmt is the "format" package
- `Println()` is "Print Line" and reserves a full line ended with a newline
- `fmt.Println("hello fmt")`

## Variable Assignment

```go
// var should be used to define a specific type
var firstString string = "First Var"
// := can only be used inside of a function, the full...
// var varName (varType) syntax must be used outside of the function
second := "Second"
fmt.Println(firstString, second)
var varName varType // must include type without a direct value
var varName = value // will infer type
```

## Constants

### `const` instead of `var`

- Immutable
- Must be known when defined

```go
// Immutable
const pi float32 = 3.14159265358979
```

## Naming Variables for Imports

```go
// This variable can not be imported:
var myLocalVar string = "Because it's lowercase"
// This variable can be imported:
var MyLocalVar string = "Because it's uppercase"
fmt.Println(myLocalVar, MyLocalVar)
```

## Types

```go
var typeInt int       // 0
var typeString string // ""
var typeFloat float64 // 0.0 (or 0)
var typeArray [5]int  // [0 0 0 0 0]
var typeBool bool     // false
fmt.Println("Empty int, string, float, int array, bool\n", typeInt, "\n", typeString, "\n", typeFloat, "\n", typeArray, "\n", typeBool)

// int - Number without decimals (-10 to 151)
var intNumber float64 = 10 / 3 // 3
fmt.Println(intNumber)

// float64 - Number with decimal places 3.3333333333333335
// float64 is heavy on memory
var floatNumber float64 = float64(10) / 3  // 3.33
defaultFloat := 1.234567890123456789123456 // Inferred type: float64

// float32
// less heavy in memory but also less precise
var smallFloat float32 = 234567890123456789123456 // float32
smallerFloat := 1.2                               // also defaults to float64
fmt.Println(floatNumber, defaultFloat, smallFloat, smallerFloat)
// string - "Hello"

// bool
// binary yes/no (true or false)
var firstBool bool = true // true
fmt.Println(firstBool)

// rune
// Single unicode character ('❤️' or 'a')
// Uses single quotes - larger than a byte, prints an int value
var firstRune rune = '€' // 8364
fmt.Println("rune: ", firstRune)
var runeToString string = string(firstRune)
fmt.Println("rune to string: ", runeToString)
defaultRune := 'b' // ends up defaulting to rune, not byte
fmt.Println("defaultByte ", defaultRune)

// byte
// Smallest type, single byte ASCII character ('a' or '9')
var firstByte byte = 'A'  // 65
var secondByte byte = 'a' // 97
fmt.Println("byte: ", firstByte, secondByte)
var firstByteString string = string(firstByte) // A
fmt.Println("Bytes to Strings: ", firstByteString)

```

## Variables

### Naming conventions

- Uses mixed case: helloWorld or HelloWorld are both good names
- lowercase for within module
- uppercase for vars that need to be `exported`

```go
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
```

## Arrays

```go
// Array Syntax
[length]type{"array", "values"}
```

### Conventions and Properties

- zero indexed
- ...
- ...

```go
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
```

## Functions

### Function structure

```go
// func name(varName varType) returnType {
// 	return varName
// }

func add(a int, b int) int {
sum := a + b
  return sum
}

func multiple_returns() [2]string {
  return [2]string{"Hello", "Array"}
}
```

### Actual Functions

```go
addResponse := add(4, 5)
fmt.Println("sum of 4 and 5", addResponse)

returnArray := multiple_returns()
fmt.Println("[2]string{\"Hello\", \"Array\"}")
fmt.Println(returnArray)
```

## Printing

```go
// fmt.Println() - creates a newline at the end
firstName := "Bob"
lastName := "Bass"
fmt.Println(firstName + " " + lastName) // Bob Bass

firstNumber := 5
secondNumber := 10
fmt.Println(firstNumber + secondNumber) // 15 (as an int)
fmt.Println("15")                       // 15 (as a string)
// Mixed types can't be used together, the types need to match
// A package can be used to convert "15" to 15
multiLine := `This is a multiline string
It can be quite long`
fmt.Println(multiLine)

```

### Format string print

- default formatter or custom, outputs formatted string to standard output
- Many of these are taken from C

```go
// fmt.Print(str)
// fmt.Printf(format, str)
// fmt.Println(str)
fmt.Print("Hello Print") // This is the previous line.Hello Print
```

### Format (Placeholders)

- %v - value
- Look up the fmt package docs to get all of the placeholder types

```go
fullName := firstName + " " + lastName
age := 30
fmt.Printf("I am %v and I am %v years old.", fullName, age) // I am Bob Bass and I am 30 years old.
* Type can be added
fmt.Printf("I am %v and I am %v (Type: %T) years old.", fullName, age, age) // I am Bob Bass and I am 30 (Type: int) years old.
fmt.Println("Hello Println")
```

## Sprint

- Format string (default formatter or custom) and return formatted string
- Returns the value as well

```go
otherFullName := fmt.Sprintf("%v %v", firstName, lastName) // Bob Bass
fmt.Println(otherFullName)

// Sprint types:
// fmt.Sprint(str)
// fmt.Sprintf(format, str)
// fmt.Sprintln(str)
```

### Fprint

- Format string (default formatter or custom) and write to specified writer
- Sprint types

```go
// Fprint types:
// fmt.Fprint(writer, str)
// fmt.Fprintf(writer, format, string)
// fmt.Fprintln(writer, str)
```

# Packages and Modules

## Modules

- A module has a unique identifier and can be distributed (e.g. library)
- Every Go project goes into a new module
- Projets can use (i.e. import from) multiple modules (custom and third-party)
- Created and managed via `go mod` commands and `go.mod` file

## Packages

- Every module contains at least one package (the "main" package)
- Multiple files can make up the same package (via package instruction - i.e. `package main`)
- A module may contain multiple packages, stored in subfolders
- Can be imported (it's ultimately always a module that is imported)

### Creating `go.mod`

```bash
# You name the module in a way that matches the folder path beginning with a github.com folder so the module matches the published path
 go mod init github.com/yourOrganization/moduleFolderPath
```

### Example go.mod

```go
module github.com/716green/academind/firstapp

go 1.16
```

### Run code in a folder with a `go.mod` file

```bash
go run .
```

## Sharing across files using packages

### This can be great for organization and readability

### There should be a subfolder with a one word name for each additional package

```go
// firstapp.go

package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// greetingText := "Hello World"
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)
}

```

### variables.go (subfolder)

```go
// Using the same package (main)
// package main

// As a separate package
// ./greeting/variables.go
package greeting

var greetingText string = "Hello World from variables.go"

```

## Importing Packages

```go
import (
	"fmt"

	"github.com/716green/academind/firstapp/greeting"
)
// or
import "fmt"
import "github.com/716green/academind/firstapp/greeting"

// Auto-format will change the second version to the first one
```
