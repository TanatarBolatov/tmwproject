package main

import "fmt"

func PrintIf() {
	age := 68

	if age < 18 {
		fmt.Println("Ты несовершеннолетний")
	} else if age >= 18 && age < 63 {
		fmt.Println("Ты совершеннолетний")
	} else {
		fmt.Println("Ты пенсионнер")
	}
}
