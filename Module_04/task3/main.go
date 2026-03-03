//— Написать две функции, которые выводят на экран любые слова любое количество раз
//— Вызвать написанные функции, миксуя их с выводами на экран в main функции
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 🔼
	// 🔲
	// ❌

	score := 0

	fmt.Println("Get Ready!")
	fmt.Println("Счёт:", score)

	for {
		fmt.Println("Я подпрыгиваю к шипам!")
		fmt.Println("----------------------")
		fmt.Println("")
		fmt.Println("🔲 🔼🔼🔼")
		fmt.Println("----------------------")
		fmt.Println("")

		fmt.Println("Я перепрыгиваю шипы!")
		fmt.Println("----------------------")
		fmt.Println("  🔲 ")
		fmt.Println("🔼🔼🔼")
		fmt.Println("----------------------")
		fmt.Println("")

		// 0 1 2 3 4 5 6 7
		if rand.Intn(8) == 1 {
			fmt.Println("Я упал в шипы(((")
			fmt.Println("----------------------")
			fmt.Println("  ❌ ")
			fmt.Println("🔼🔼🔼")
			fmt.Println("----------------------")
			fmt.Println("")

			break
		}

		fmt.Println("Я успешно пролетел через шипы!")
		fmt.Println("----------------------")
		fmt.Println("")
		fmt.Println("🔼🔼🔼 🔲")
		fmt.Println("----------------------")
		fmt.Println("")

		score += 5
		fmt.Println("Счёт:", score)
		fmt.Println("")
		fmt.Println("==============================")
		fmt.Println("")

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Конец игры!")
	fmt.Println("Итоговый счёт:", score)
}
