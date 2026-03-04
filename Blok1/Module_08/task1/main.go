//— Придумать себе любую структуру, неважно пользователь, автомобиль или питомец
//— Описать для этой структуры поля, которые посчитаете нужными
//— Создать несколько экземпляров описанной структуры, вывести их на экран
//— Поменять значения полей структуры после создания, выводя на экран значения поля до и после изменения
package main

import "fmt"

type Car struct {
	HorsePower int
	Marka      string
	Model      string
	Class      string
}

func main() {
	car1 := Car{
		HorsePower: 150,
		Marka:      "Haval",
		Model:      "Jolion",
		Class:      "Кроссовер",
	}

	car2 := Car{
		HorsePower: 550,
		Marka:      "BMW",
		Model:      "X5",
		Class:      "Седан",
	}

	fmt.Println("car1:", car1)
	fmt.Println("car2:", car2)

	fmt.Println("Лошадинные силы до чип-тюнинга:", car1.HorsePower)
	car1.HorsePower += 100
	fmt.Println("Лошадинные силы после чип-тюнинга:", car1.HorsePower)
}
