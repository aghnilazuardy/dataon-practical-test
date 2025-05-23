package main

import (
	"fmt"
	"strconv"
)

func numberBreakdown(number int) []int {
	var result []int

	numberString := strconv.Itoa(number)
	numberStringLength := len(numberString)

	for i, digitChar := range numberString {
		digit, _ := strconv.Atoi(string(digitChar))

		if digit == 0 {
			continue
		}

		numberPlace := 1
		for j := 0; j < numberStringLength-i-1; j++ {
			numberPlace *= 10
		}

		numComponent := digit * numberPlace
		result = append(result, numComponent)
	}

	return result
}

func main() {
	number := 1225441
	fmt.Printf("Angka: %v\n", number)
	fmt.Println("Output:")

	components := numberBreakdown(number)
	for _, comp := range components {
		fmt.Println(comp)
	}
}
