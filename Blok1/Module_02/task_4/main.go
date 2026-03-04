package main

import "fmt"

func main() {
	number := 50
	score := 10

	fmt.Println("number до +:", number)
	number = score + 15
	fmt.Println("number после +:", number)

	fmt.Println("number до -:", number)
	number -= 25
	fmt.Println("number после -:", number)

	fmt.Println("number до *:", number)
	number *= 3
	fmt.Println("number после *:", number)

	fmt.Println("number до /:", number)
	number /= 5
	fmt.Println("number после /:", number)

	fmt.Println("number до %:", number)
	number %= 5
	fmt.Println("number после %:", number)

	i := 10
	fmt.Println("i до инкремента:", i)
	i++
	fmt.Println("i после инкремента:", i)

	fmt.Println("i до декремента:", i)
	i--
	fmt.Println("i после декремента:", i)

	text := "hello"
	fmt.Println("text до +:", text)
	text = text + "world"
	fmt.Println("text после +:", text)
}
//— Создать переменные всех изученных типов (int, string, float64, bool) при помощи любого вида записи
//— Произвести над созданными переменными все изученные операции (+, -, *, /, %, +=, -=, *=, /=, %=, ++, --), каждый раз выводя переменную на экран до произведения операции и после