package main

import (
	"fmt"

	"gitlab.com/Kdrai007/packages-tryout/mystring"
)

func main() {
	fmt.Println("----Packages and Modules in Go-----")
	str := mystring.Reverse("something in it")
	fmt.Println(str)
}
