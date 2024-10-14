package main

import (
	"fmt"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	height, width float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func printShape(s shape) {
	fmt.Printf("area = %v\n", s.area())
	fmt.Printf("perimeter= %v", s.perimeter())
}

func checkSwitch(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Println("This is an integer", v)
	case string:
		fmt.Println("This is a string", v)
	default:
		fmt.Println("This is another type", v)
	}
}

func main() {
	// ar := rect{width: 24, height: 25}
	// printShape(ar)
	fmt.Println("Let's learn about interfaces in Go!")

	// exerciseOne()
	// exerciseTwo()
	// exerciseThree()
	// exerciseFour()
	exerciseFive()

	//------------- type-assertion---------------

	// var i interface{} = "hello"
	// c := i.(string)

	// fmt.Println(c) // output: hello

	// s := shape(rect{width: 24, height: 25})

	// c, ok := s.(rect)

	// if ok {
	// 	fmt.Println(c)
	// } else {
	// 	fmt.Println("Type assertion failed")
	// }

	//------------- type-switch ---------------

	// checkSwitch("hello there")
	// checkSwitch(10)
	// checkSwitch(20.25)
}
