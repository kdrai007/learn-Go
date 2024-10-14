package main

import "fmt"

func bulkSend(totalMsg int) float64 {
	var totalCost float64

	for i := 0; i < totalMsg; i++ {
		baseCost := 1.0
		fee := float64(i) * 0.01
		totalCost += baseCost + fee
	}

	return totalCost
}

func maxMessages(thresh int) int {
	totalCost := 0
	for i := 0; ; i++ {
		totalCost += 100 + i
		if totalCost >= thresh {
			return i
		}
	}
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance >= 0 {
		actualCostInPennies *= costMultiplier

		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func main() {
	// fmt.Printf("bulk cost would be %.2f", bulkSend(12))
	// fmt.Println(maxMessages(1000))

	fmt.Println(getMaxMessagesToSend(1.3, 13))
}
