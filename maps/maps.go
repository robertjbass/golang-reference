package main

import "fmt"

func newSection(title string) {
	const divider = "----------------------"
	fmt.Printf("%v%v%v", "\n", divider, "\n")
	fmt.Println(title)
	fmt.Println(divider)
}

func main() {
	// websites := []string{"https://www.google.com", "https://aws.com"}
	//* Struct shapes are immutable, maps are not
	emptyMap := map[string]string{}
	fmt.Println(emptyMap)
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://www.aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	fmt.Println(websites["Amazon Web Services"])

	newSection("Add Map Value")

	websites["Azure"] = "https://azure.microsoft.com"
	fmt.Println(websites["Azure"])

	newSection("Delete Map Value")

	delete(websites, "Google")
	fmt.Println(websites)

	// Length of map
	fmt.Println(len(websites))

}
