package main

import (
	"fmt"
	"sort"
)

type MoneyBreakdown struct {
	Denomination int
	Count        int
	TotalValue   int
}

func calculateMoneyBreakdown(amount int, denominations []int) []MoneyBreakdown {
	var result []MoneyBreakdown

	// Urutkan denominasi dari yang terbesar ke terkecil
	sort.Sort(sort.Reverse(sort.IntSlice(denominations)))

	remaining := amount

	for _, denom := range denominations {
		if remaining >= denom {
			count := remaining / denom
			totalValue := count * denom

			result = append(result, MoneyBreakdown{
				Denomination: denom,
				Count:        count,
				TotalValue:   totalValue,
			})

			remaining -= totalValue
		}
	}

	return result
}

func main() {
	targetAmount := 1575250
	availableDenominations := []int{100000, 50000, 20000, 5000, 100, 50}

	fmt.Printf("Jumlah uang yang akan diambil: %v\n", targetAmount)
	fmt.Println("\nPecahan uang yang tersedia:")
	for _, denom := range availableDenominations {
		fmt.Printf("- %v\n", denom)
	}

	breakdowns := calculateMoneyBreakdown(targetAmount, availableDenominations)

	fmt.Println("\nHasil Pecahan Uang:")
	for _, b := range breakdowns {
		fmt.Printf("- %v: %d lembar\n", b.Denomination, b.Count)
	}

	totalSheets := 0
	for _, b := range breakdowns {
		totalSheets += b.Count
	}
	fmt.Printf("\nTotal lembar uang yang dibutuhkan: %d lembar\n", totalSheets)

}
