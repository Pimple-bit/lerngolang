//— Написать две функции, которые выводят на экран любые слова любое количество раз
//— Вызвать написанные функции, миксуя их с выводами на экран в main функции
package main

import "fmt"

// Я main и я начался!
// 1
// 2
// 3
// Я main 1
// Я main 2
// Я main 3
// 11
// 22
// 1
// 2
// 3
// 33
// Я main и я закончился.

func main() {
	fmt.Println("Я main и я начался!")
	foo()
	fmt.Println("Я main 1")
	fmt.Println("Я main 2")
	fmt.Println("Я main 3")
	boo()
	fmt.Println("Я main и я закончился.")
}

func foo() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

func boo() {
	fmt.Println("11")
	fmt.Println("22")
	foo()
	fmt.Println("33")
}
