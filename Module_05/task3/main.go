// /— База обучения по теме "Функции" в программировании:
//— Напишите функцию, которая принимает коэффициенты 'a', 'b', 'c' квадратного уравнения, и находит корни этого квадратного уравнения
//— Найденные корни необходимо вывести на экран
package main

import (
	"fmt"
	"math"
)

func main() {
	a := 9
	b := 6
	c := 1

	square(a, b, c)
}

func square(a int, b int, c int) {
	d := b*b - 4*a*c

	if d < 0 {
		fmt.Println("Корней нет!")
		return
	}

	if d == 0 {
		x := float64(-b) / float64(2*a)

		fmt.Println("У уравнения один корень:", x)

		return
	}

	// int + float64

	x1 := (float64(-b) + math.Sqrt(float64(d))) / float64(2*a)
	x2 := (float64(-b) - math.Sqrt(float64(d))) / float64(2*a)

	fmt.Println("Первый корень:", x1)
	fmt.Println("Второй корень:", x2)
}
