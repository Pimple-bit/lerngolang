//— Создать несколько функций, помимо main функции
//— В каждой созданной функции через отложенный вызов необходимо выводить на экран любой текст
//— Помимо выводов на экран в отложенных функциях, необходимо выводить на экран непосредственно в описанных рядом с main функциях
//— Вызвать описанные функции в main функции, параллельно совершая в main функции обычные выводы на экран
//— Попробовать предугадать вывод программы
//— Запустить программу и сравнить ожидания с реальностью
package main

import "fmt"

// "Я main и я начался"
// "Вызов foo()"

// "foo 1"
// "foo 2"
// "defer foo 2"
// "defer foo 1"

// "Вызов boo()"

// "boo 1"
// "boo 2"
// "defer boo 2"
// "defer boo 1"

// "Я main и я закончился"

func main() {
	fmt.Println("Я main и я начался")
	defer func() {
		fmt.Println("Я main и я закончился")
	}()

	fmt.Println("Вызов foo()")
	foo()
	fmt.Println("Вызов boo()")
	boo()
}

func foo() {
	defer func() {
		fmt.Println("defer foo 1")
	}()

	fmt.Println("foo 1")

	defer func() {
		fmt.Println("defer foo 2")
	}()

	fmt.Println("foo 2")
}

func boo() {
	fmt.Println("boo 1")

	defer func() {
		fmt.Println("defer boo 1")
	}()

	fmt.Println("boo 2")

	defer func() {
		fmt.Println("defer boo 2")
	}()
}
