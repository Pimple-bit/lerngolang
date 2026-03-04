// — Написать любой вывод на экран в отложенных анонимных функциях, сами функции "задефёрить" в main функции
// — Помимо выводов на экран в отложенных функциях, необходимо вывести на экран любые тексты в main функции
// — Попробовать предугадать вывод программы
// — Запустить программу и сравнить ожидания с реальностью
package main

import "fmt"

// "main 1"
// "main 2"
// "main 3"
// "defer 3"
// "defer 2"
// "defer 1"

// стек отложенных вызовов
//
//
//
//

//

func main() {
	defer func() {
		fmt.Println("defer 1")
		}()

	fmt.Println("main 1")

	defer func() {
		fmt.Println("defer 2")
	}()

	fmt.Println("main 2")

	defer func() {
		fmt.Println("defer 3")
	}()

	fmt.Println("main 3")
}
