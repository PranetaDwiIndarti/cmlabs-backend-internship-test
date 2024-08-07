package main

import (
	"fmt"
)

func fizzBuzzTest(firstNumber, lastNumber int) {
	fmt.Println("FizzBuzz with Golang")
	for i := firstNumber; i <= lastNumber; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzBuzzTest(1, 100)
}


