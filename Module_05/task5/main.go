//— Создать глобальную переменную
//— Описать несколько разных функций, которые бы меняли эту глобальную переменную
// /— Вызвать описанные фукнции, и посмотреть на все промежуточные и на итоговое значение созданной глобальной переменной
package main

import "fmt"

var drob float64 = 10.55

func main() {
	fmt.Println("drob 1:", drob)
	foo()
	fmt.Println("drob 2:", drob)
	boo()
	fmt.Println("drob 3:", drob)
	goo()
	fmt.Println("drob 4:", drob)
}

func foo() {
	drob *= 2
}

func boo() {
	drob += 1.25
}

func goo() {
	drob /= 3
}
