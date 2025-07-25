package main

import "fmt"

func LessonRev() {
	mass := []int{0, 1, 2, 3, 4, 5}

	for a, b := 0, len(mass)-1; a < b; a, b = a+1, b-1 {
		mass[a], mass[b] = mass[b], mass[a]
	}
	fmt.Println(mass)
}
