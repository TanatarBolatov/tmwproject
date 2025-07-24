package main

import "fmt"

func PrintInfo() {
	var age int = 33
	var height string = "Постоянный"
	name := "Tanatar"
	isStudent := true

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Рост:", height)
	fmt.Println("Студент?", isStudent)
}
