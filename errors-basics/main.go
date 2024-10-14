package main

import (
	"fmt"
	"strconv"
)

// type error interface {
// 	Error() string
// }

func Atoi(s string) {
	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Printf("couldn't convert\n %v \n", err)

		return
	}

	fmt.Println("value of i: ", i)
	fmt.Println("value of string: ", s)
}

func subStr() string {
	// value := 111
	// s := fmt.Sprintf("%v means! run like your life defends on it\n", value)
	// fmt.Printf(s)

	//Rouding off floats

	cost := 200
	recipient := "someone"
	s := fmt.Sprintf("SMS that costs %v to be sent to '%v' can not be sent", cost, recipient)
	return s
}

func main() {
	// exerciseOne()
	exerciseTwo()
	// fmt.Println(subStr())
}
