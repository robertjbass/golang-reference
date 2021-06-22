package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	const url string = "https://reqres.in/api/users?page=2"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)

}
