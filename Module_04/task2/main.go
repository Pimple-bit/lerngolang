// /— Аналогично программам из видео-урока, напишите продвинутую симуляцию игры в Geometry Dash
//— В симуляции должны быть бесконечные уровни
//— Необходимо предусмотреть возможность проигрыша
//— За каждое преодолённое препятствие начисляется 5 очков
//— В конце симуляции необходимо вывести на экран количество заработанных очков
package main

import (
	"fmt"
	"time"
)

func main() {
	// 🔼
	// 🔲
	// ❌

	score := 0

	fmt.Println("Get Ready!")
	fmt.Println("Счёт:", score)

	for i := 1; i <= 3; i++ {
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

		time.Sleep(1 * time.Second)
	}

	fmt.Println("Конец игры!")
	fmt.Println("Итоговый счёт:", score)

}
