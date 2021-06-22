package main

import (
	"fmt"

	"github.com/716green/academind/bmi-app/convert"
	"github.com/716green/academind/bmi-app/info"
)

func main() {

	fmt.Println(info.MainTitle)
	fmt.Println(info.Seperator)

	fmt.Println(info.FtPrompt)
	fmt.Print(info.FtPrompt2)

	//* '\n' listens for the enter key
	inputFeet, _ := reader.ReadString('\n')
	inputFeet = convert.Trim(inputFeet)

	fmt.Printf("%vft →__in: ", inputFeet)
	inputInches, _ := reader.ReadString('\n')
	inputInches = convert.Trim(inputInches)
	fmt.Printf("You entered %v' %v\"\n", inputFeet, inputInches)
	meters := convert.ConvertFtIntoM(inputFeet, inputInches)

	fmt.Print(info.LbsPrompt)
	inputWeight, _ := reader.ReadString('\n')
	inputWeight = convert.Trim(inputWeight)
	kgWeight := convert.LbsToKg(inputWeight)

	bmi := convert.ConvertBmiFromKgAndM(meters, kgWeight)
	fmt.Printf("Your BMI is %.2f\n", bmi)

}
