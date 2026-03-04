// — Мы описываем операцию по изменению размера груди
// — Пусть эта операция описана функцией, которая может изменить значение переменной, созданной в main-функции
// — Размер груди задан дробной переменной в main'е
// — Описать функцию изменения размера груди, протестировать её
package main

import "fmt"

func main() {
	size := 2.5

	var ptr *float64 = nil

	fmt.Println("размер до:", size)
	changeSize(ptr)
	fmt.Println("размер после:", size)
}

func changeSize(sizePtr *float64) {
	if sizePtr == nil {
		fmt.Println("Переданный указатель указывает вникуда!")
		fmt.Println("Операция по изменению размера груди отменена.")
		return
	}

	*sizePtr += 1.0
}
