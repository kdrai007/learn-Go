package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgArr := [3]string{primary, secondary, tertiary}
	costArr := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}

	return msgArr, costArr
}

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr[:])

	exerciseOne()
}
