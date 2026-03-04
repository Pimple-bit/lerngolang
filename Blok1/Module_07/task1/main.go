//— Создайте указатели на переменные всех изученных типов (int, string, bool, float64)
//— Выведите на экран адреса в памяти, которые в себе хранят эти указатели
//— Выведите на экран значения переменных, на которые указатели указывают, при помощи разыменования указателей
package main

import "fmt"

func main() {
	number := 15
	text := "hello"
	boolean := false
	drob := 56.76

	numberPtr := &number
	textPtr := &text
	booleanPtr := &boolean
	drobPtr := &drob

	fmt.Println("numberPtr:", numberPtr)
	fmt.Println("textPtr:", textPtr)
	fmt.Println("booleanPtr:", booleanPtr)
	fmt.Println("drobPtr:", drobPtr)

	fmt.Println("number:", *numberPtr)
	fmt.Println("text:", *textPtr)
	fmt.Println("boolean:", *booleanPtr)
	fmt.Println("drob:", *drobPtr)
}
