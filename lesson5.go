package main

import "fmt"

func Lesson5() {
	mass := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Массивы:", mass)
	s := []string{"a", "b", "c"}
	fmt.Println("Слайсы до:", s, "len=", len(s), "cap=", cap(s))
	s = append(s, "d", "e")
	fmt.Println("Слайсы после append:", s, "len=", len(s), "cap=", cap(s))
}
