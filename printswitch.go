package main

import "fmt"

func PrintSwitch() {
	role := "admin"

	switch role {
	case "user":
		fmt.Println("Привет пользователь!")
	case "admin":
		fmt.Println("Здравствуйте администратор!")
	default:
		fmt.Println("Неизвестная роль!")
	}
}
