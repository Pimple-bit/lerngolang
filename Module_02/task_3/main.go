package main

import "fmt"

func main() {
	number1 := 15
	text1 := "hello"
	drob1 := 5.65
	boolean1 := false

	var number2 int = 25
	var text2 string = "world"
	var drob2 float64 = 6.775
	var boolean2 = true

	var number3 int
	var text3 string
	var drob3 float64
	var boolean3 bool

	fmt.Println("number1:", number1)
	fmt.Println("text1:", text1)
	fmt.Println("drob1:", drob1)
	fmt.Println("boolean1:", boolean1)

	fmt.Println("--------------------")

	fmt.Println("number2:", number2)
	fmt.Println("text2:", text2)
	fmt.Println("drob2:", drob2)
	fmt.Println("boolean2:", boolean2)

	fmt.Println("--------------------")

	fmt.Println("number3:", number3)
	fmt.Println("text3:", text3)
	fmt.Println("drob3:", drob3)
	fmt.Println("boolean3:", boolean3)
}
//— Создать переменные всех изученных типов (int, string, float64, bool) при помощи короткого вида записи
// //— Создать переменные всех изученных типов (int, string, float64, bool) при помощи длинного вида записи и присвоить им значения
// //— Создать переменные всех изученных типов (int, string, float64, bool) при помощи длинного вида записи и НЕ присваивать им никакого значения (оставить значения по-умолчанию)
// //— Вывести все созданные переменные на экран. Перед каждой переменной, при выводе на экран, необходимо указать её название.