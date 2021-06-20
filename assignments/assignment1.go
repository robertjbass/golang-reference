//* 1 - Create a new Go code file and initialize it correctly
//* 2 - Add two new variables to the file, (PI and A circle with a radius value of 5)
//* 3 - Calculate the circle circumference (2*PI*radius) and store it in a new variable.
//? Output that value in the command line. Make sure that the value is correct.
//* 4 - Switch to a different way of outputting the results:
//? Format the string to say "For a radius of X, the circle circumfrence is Y.YY".
//? X and Y should be concrete values.
//! Hint: Look into https://golang.org/pkg/fmt/#hdr-Printing to learn how to output 2 decimal places

package main

import (
	"fmt"
)

func main() {
	var pi float64 = 3.14159265358979
	radius := 5
	circleCircumference := calculateCircumference(pi, radius)
	fmt.Printf("For a radius of %v the circumference is %v\n", radius, circleCircumference)
	testValue := (5 * 2 * 3.14159265358979)
	fmt.Printf("Test Value: %v", 5*2*3.14159265358979)
	valuesMatch := testValue == circleCircumference
	if valuesMatch {
		fmt.Println("Values Match")
	} else {
		fmt.Println("Values Don't Match")
	}
}

func calculateCircumference(pi float64, radius int) float64 {
	circumference := 2.0 * pi * float64(radius)
	return circumference
}
