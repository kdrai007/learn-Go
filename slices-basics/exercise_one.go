package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		return messages[:], nil
	}
	if plan == planFree {
		return messages[:2], nil
	}
	return nil, errors.New("Unsupported plan")
}

func exerciseOne() {
	fmt.Println("--------exercise One----------")
	arr, err := getMessageWithRetriesForPlan("pfree", [3]string{"message1", "message2", "message3"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr)

}
