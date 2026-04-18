package main

import (
	"fmt"
	"math"
)


type Shape interface {
	Area() float64
	Perimeter() float64
}


type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}


type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}


func PrintShapeInfo(s Shape) {
	fmt.Println("Площадь:", s.Area())
	fmt.Println("Периметр:", s.Perimeter())
	fmt.Println("------")
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{A: 3, B: 4, C: 5}

	PrintShapeInfo(rect)
	PrintShapeInfo(circle)
	PrintShapeInfo(triangle)
}
//— Необходимо разработать модуль для работы с геометрическими фигурами
//— Каждая фигура умеет вычислять свою площадь и периметр
//— Разные фигуры реализуются по-разному, но взаимодействие с ними должно быть унифицировано через интерфейс
//— Необходимо описать фигуры и функцию, которая сможет принимать различные виды фигур и высчитывать их площать и периметр