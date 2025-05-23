package main

import (
	"fmt"
	"regexp"
)

func main() {
	email := "user@example.com"

	regex := regexp.MustCompile(`@([\w.-]+\.[a-zA-Z]{2,})`)

	result := regex.FindStringSubmatch(email)

	if len(result) > 1 {
		fmt.Println(result[1])
	} else {
		fmt.Println("No match")
	}
}
