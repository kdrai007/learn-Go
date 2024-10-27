package main

import (
	"fmt"
	"strings"
)

func replaceBadWords(message *string) {
	cussWords := map[string]string{
		"fuck": "****",
		"suck": "****",
	}
	for word, replacement := range cussWords {
		*message = strings.ReplaceAll(*message, word, replacement)
	}
}

func main() {
	fmt.Println("----Pointer basic------")

	// message := ""
	// replaceBadWords(&message)
	// fmt.Println(message)
	// exerciseOne()
	exerciseTwo()
}
