package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for i := 2; i < max; i++ {
		if i == 2 {
			fmt.Println("prime number: ", i)
			continue
		}
		if i%2 == 0 {
			continue
		}
		isPrime := true
		for j := 3; j <= int(math.Sqrt(float64(i))); j += 2 {
			if i%j == 0 {
				isPrime = true

			}
		}
		if isPrime {
			fmt.Println("prime number: ", i)
		}
	}

}

func exerciseTwo() {
	printPrimes(45)
}
