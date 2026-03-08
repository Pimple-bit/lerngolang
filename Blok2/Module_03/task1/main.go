package main

import "fmt"

func catchPanic() {
	if r := recover(); r != nil {
		fmt.Println("Поймали панику:", r)
	}
}

func divideByZero() {
	defer catchPanic()

	a := 10
	b := 0

	fmt.Println("Деление на ноль:")
	fmt.Println(a / b)
}

func writeToNilMap() {
	defer catchPanic()

	fmt.Println("Запись в nil map:")

	var m map[string]int
	m["test"] = 1
}

func outOfSlice() {
	defer catchPanic()

	fmt.Println("Выход за пределы слайса:")

	s := []int{1, 2, 3}
	fmt.Println(s[3])
}

func manualPanic() {
	defer catchPanic()

	fmt.Println("Ручная паника:")
	panic("это наша собственная паника")
}

func main() {

	divideByZero()
	fmt.Println()

	writeToNilMap()
	fmt.Println()

	outOfSlice()
	fmt.Println()

	manualPanic()
}