//— Придумать ограничения на созданные поля структуры
//— Для описанной в предыдущем задании структуры необходимо создать несколько методов, используя которые можно бы было поддержвать значения полей структуры в согласованном состоянии (прикрутить валидацию), обратить внимание где лучше использовать ресивер по значению, а где по указателю
//— Для описанной структуры создать конструктор с валидацией передаваемых значений
//— Протестировать всю написанную функциональность, как можно больше выводов на экран, чтобы понимать что происходит в каждой отдельной строчке кода, какие методы/функции вызываются/какие значения полям присваиваются и как эти самые поля изменяются
package main

import "fmt"

type Car struct {
	// HorsePower >= 50
	// HorsePower <= 1000
	HorsePower int

	Marka string
	Model string
	Class string

	// RazgonDoSta <= 20.0s
	// RazgonDoSta >= 1.0s
	RazgonDoSta float64
}

func NewCar(
	horsePower int,
	marka string,
	model string,
	class string,
	razgonDoSta float64,
) Car {
	if horsePower < 50 || horsePower > 1000 {
		fmt.Println("horsePower не прошла валидацию!")
		return Car{}
	}

	if marka == "" {
		fmt.Println("marka не прошла валидацию!")
		return Car{}
	}

	if model == "" {
		fmt.Println("model не прошла валидацию!")
		return Car{}
	}

	if class == "" {
		fmt.Println("class не прошёл валидацию!")
		return Car{}
	}

	if razgonDoSta < 1 || razgonDoSta > 20 {
		fmt.Println("razgonDoSta не прошёл валидацию!")
		return Car{}
	}

	return Car{
		HorsePower:  horsePower,
		Marka:       marka,
		Model:       model,
		Class:       class,
		RazgonDoSta: razgonDoSta,
	}
}

func (c *Car) HorsePowerUp(horsePower int) {


	if c.HorsePower+horsePower > 1000 {
		fmt.Println("Вы не можете на столько сильно увеличить мощность машины")
		return
	}

	c.HorsePower += horsePower
}

func (c *Car) HorsePowerDown(horsePower int) {


	if c.HorsePower-horsePower < 50 {
		fmt.Println("Вы не можете на столько сильно уменьшить мощность машины")
		return
	}

	c.HorsePower -= horsePower
}

func (c *Car) ChangeMarka(newMarka string) {
	if newMarka == "" {
		fmt.Println("Марка машины не может быть пустой")
		return
	}

	c.Marka = newMarka
}

func (c *Car) ChangeModel(newModel string) {
	if newModel == "" {
		fmt.Println("Модель машины не может быть пустой")
		return
	}

	c.Model = newModel
}

func (c *Car) ChangeClass(newClass string) {
	if newClass == "" {
		fmt.Println("Класс машины не может быть пустым")
		return
	}

	c.Class = newClass
}

func (c *Car) RazgonDoStaUp(razgonDoSta float64) {
	if c.RazgonDoSta+razgonDoSta > 20 {
		fmt.Println("Вы не можете на столько сильно замедлить вашу машину")
		return
	}

	c.RazgonDoSta += razgonDoSta
}

func (c *Car) RazgonDoStaDown(razgonDoSta float64) {
	if c.RazgonDoSta-razgonDoSta < 1 {
		fmt.Println("Вы не можете на столько сильно ускорить вашу машину")
		return
	}

	c.RazgonDoSta -= razgonDoSta
}

func main() {
	car := NewCar(550, "Honda", "Civic", "Седан", 10.2)

	fmt.Println("car до:", car)
	car.HorsePowerUp(500)

	fmt.Println("car после:", car)

}
