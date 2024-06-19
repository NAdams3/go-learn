package main

import "fmt"

func FizzBuzzPrinter(n int) {

	if n%3 == 0 {
		fmt.Print("Fizz")
	}

	if n%5 == 0 {
		fmt.Print("Buzz")
	}

	if n%5 != 0 && n%3 != 0 {
		fmt.Printf("%v", n)
	}

	fmt.Println()

}
