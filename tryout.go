package main

import "fmt"

func main() {
	x := 50
	y := &x
	*y = 100
	fmt.Println(*y)
	fmt.Println(x)

	var points *int = &x

	// checking if pointer value is not nil

	if points != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(*points)
}
