package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgArr := [3]string{primary, secondary, tertiary}
	costArr := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}

	return msgArr, costArr
}

func getMessageCost(messages []string) []float64 {
	messagesCost := make([]float64, len(messages))
	for i, msg := range messages {
		messagesCost[i] = float64(len(msg)) * 0.01
	}
	return messagesCost
}

func aboutAppend(slice []int) {
	arr := []int{1, 2, 3, 4, 5}
	slice = append(slice, arr...)
	fmt.Println(slice)
}

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// mySlice := arr[:5]
	// mySlice[4] = 100
	// fmt.Println(cap(mySlice))
	// aboutAppend([]int{3, 6, 7})

	// exerciseOne()
	// exerciseTwo()
	// exerciseThree()
	exerciseFour()

	// fmt.Println(getMessageCost([]string{"Welcome to the movies!", "Enjoy your popcorn!"}))
}
