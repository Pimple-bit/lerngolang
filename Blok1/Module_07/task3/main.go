//— Создайте указатели на переменные всех изученных типов (int, string, bool, float64)
//— Для каждого указателя создайте ему в пару второй указатель с nil-значением
//— Опишите для каждого указателя функцию, которая бы принимала указатель, выводила его на экран, а затем проверяла является ли указатель nil-указателем
//— Если указатель nil, то выведите на экран сообщение об этом
//— Если указатель не nil, то выведите значение переменной, которая лежит под этим указателем
package main

import "fmt"

func main() {
	// number := 10
	// text := "hello"
	// boolean := true
	// drob := 5.45

	// numberPtr := &number
	// textPtr := &text
	// booleanPtr := &boolean
	// drobPtr := &drob

	var numberPtr2 *int = nil
	// var textPtr2 *string = nil
	// var booleanPtr2 *bool = nil
	// var drobPtr2 *float64 = nil

	fooInt(numberPtr2)

}

func fooInt(ptr *int) {
	fmt.Println("int ptr:", ptr)

	if ptr == nil {
		fmt.Println("int ptr никуда не указывает!")
		return
	}

	fmt.Println("int ptr указывает на переменную со значением:", *ptr)
}

func fooString(ptr *string) {
	fmt.Println("string ptr:", ptr)

	if ptr == nil {
		fmt.Println("string ptr никуда не указывает!")
		return
	}

	fmt.Println("string ptr указывает на переменную со значением:", *ptr)
}

func fooBoolean(ptr *bool) {
	fmt.Println("bool ptr:", ptr)

	if ptr == nil {
		fmt.Println("bool ptr никуда не указывает!")
		return
	}

	fmt.Println("bool ptr указывает на переменную со значением:", *ptr)
}

func fooFloat64(ptr *float64) {
	fmt.Println("float64 ptr:", ptr)

	if ptr == nil {
		fmt.Println("float64 ptr никуда не указывает!")
		return
	}

	fmt.Println("float64 ptr указывает на переменную со значением:", *ptr)
}
