package main

import (
	"errors"
	"fmt"
	"math"
)

func swap(str1, str2 string) (string, string, error) {
	if str1 == "" {
		return str1, str2, errors.New("something is wrong here")
	}
	return str1, str2, nil
}
func powerOf(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

var i, j int = 1, 2

const pi = 3.14159

func main() {
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	//	var c, python, java = false, false, "no!"
	//c, python, java := false, false, "no!"
	//fmt.Println(i, j, c, python, java)
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum = sum + i
	//}
	//fmt.Println(sum)
	//fmt.Println(powerOf(2, 4, 10), powerOf(2, 2, 5))
	//fmt.Println(i, j)
	//fmt.Println("value of pi is: ", pi)

	// Runes and String Encoding
	//const name = "🐻"
	//fmt.Println(name)
	//var boolean bool = true
	//if boolean {
	//fmt.Println("hello there! you are working pretty fine")
	//}

	var email string = "kundan@man.com"
	for i, r := range email {
		if r == '@' {
			user := email[:i]
			domain := email[i+1:]
			fmt.Println(user)
			fmt.Println(domain)
			break
		}
	}
}
