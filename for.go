package main

import "fmt"

func For() {
	for i := 0; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		if i%3 != 0 || i%5 != 0 {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
