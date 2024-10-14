package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	cost, err := sendSMS(msgToCustomer)
	totalCost := 0
	if err != nil {
		return 0, err
	}
	totalCost += cost
	cost, err = sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	totalCost += cost
	return totalCost, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func exerciseOne() {
	totalCost, err := sendSMSToCouple("Hello, Customer!", "Hello, Spouse!")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Total cost:", totalCost)
	}

}
