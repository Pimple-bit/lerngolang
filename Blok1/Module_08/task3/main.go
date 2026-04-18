//— Мы описываем квартиру
//— У каждой квартиры есть адрес, площадь, этаж, и стоимость
//— Каждая квартира может менять свою стоимость
//— Адрес не должен быть пустой, и не может изменяться
//— Площадь должна быть больше 10 квадратных метров и не может изменяться
//— Этаж должен быть больше нуля и меньше либо равен ста
//— Стоимость должна быть больше нуля и может изменяться
//— Необходимо описать соответсвующую структуру с конструктором, полями и методами
//— То, что поля на текущем этапе изучение мы сможем изменить напрямую и сделать их невалидными мы во внимание не берём
package main

import "fmt"

type Flat struct {
	Address string
	Space   float64
	Level   int
	Cost    int 
}

func NewFlat(address string, space float64, level int, cost int) Flat {
	if address == "" {
		fmt.Println("address не прошёл валидацию!")
		return Flat{}
	}

	if space <= 10.0 {
		fmt.Println("space не прошла валидацию")
		return Flat{}
	}

	if level <= 0 || level > 100 {
		fmt.Println("level не прошёл валидацию!")
		return Flat{}
	}

	if cost <= 0 {
		fmt.Println("cost не прошла валидацию!")
		return Flat{}
	}

	return Flat{
		Address: address,
		Space:   space,
		Level:   level,
		Cost:    cost,
	}
}

func (f *Flat) ChangeCost(cost int) {
	if f.Cost+cost <= 0 {
		fmt.Println("Вы не можете задать НЕ положительную стоимость!")
		return
	}

	f.Cost += cost
}

func main() {
	flat := NewFlat("г. Нью-Йорк", 100.45, 10, 500)

	fmt.Println("flat до:", flat)
	flat.ChangeCost(-1000)
	fmt.Println("flat после:", flat)
}
