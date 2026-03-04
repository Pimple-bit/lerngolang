package main

import "fmt"

func main() {
	m1 := map[int]string{}
	m1[1] = "hi"
	m1[5] = "sun"
	m1[-103] = "cloud"

	delete(m1, 111)

	fmt.Println("m1:", m1)

	m2 := make(map[string]bool)
	m2["Петя"] = true
	m2["Вася"] = false
	m2["Ваня"] = true

	delete(m2, "Владимир")

	fmt.Println("m2:", m2)
}
