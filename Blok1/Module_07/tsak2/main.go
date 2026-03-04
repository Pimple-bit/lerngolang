//— Создайте указатели на переменные всех изученных типов (int, string, bool, float64)
//— Для каждого указатели опишите функцию, которая бы принимала этот указатель и меняла значение переменной, на которую указывает переданный указатель
//— Выведите значения переменных до вызова функций и после
package main

import "fmt"

func main() {
	number := 10
	text := "hello"
	boolean := true
	drob := 5.45

	fmt.Println("number до:", number)
	fooInt(&number)
	fmt.Println("number после:", number)

	fmt.Println("-----------------")

	fmt.Println("text до:", text)
	fooString(&text)
	fmt.Println("text после:", text)

	fmt.Println("-----------------")

	fmt.Println("boolean до:", boolean)
	fooBoolean(&boolean)
	fmt.Println("boolean после:", boolean)

	fmt.Println("-----------------")

	fmt.Println("drob до:", drob)
	fooFloat64(&drob)
	fmt.Println("drob после:", drob)
}

func fooInt(ptr *int) {
	*ptr = 15
}

func fooString(ptr *string) {
	*ptr = "world"
}

func fooBoolean(ptr *bool) {
	*ptr = false
}

func fooFloat64(ptr *float64) {
	*ptr = 60.342
}
