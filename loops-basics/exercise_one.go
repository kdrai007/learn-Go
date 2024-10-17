package main

import "fmt"

func countConnections(groupSize int) int {
	totalConnection := 0
	for i := 0; i <= groupSize; i++ {
		totalConnection += i
	}
	return totalConnection
}

func exerciseOne() {
	fmt.Println("---------Exercise One----------")
	fmt.Printf("\n\n")
	fmt.Println(countConnections(10))
}
