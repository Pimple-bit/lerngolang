package main

import "fmt"

type MyStruct struct {
	Number int
	String string
}

func main() {
	intSlice := []int{}
	intSlice = append(intSlice, 10)
	fmt.Println("intSlice:", intSlice, len(intSlice), cap(intSlice))

	strSlice := make([]string, 0)
	strSlice = append(strSlice, "hi")
	fmt.Println("strSlice:", strSlice, len(strSlice), cap(strSlice))

	floatSlice := make([]float64, 10)
	floatSlice = append(floatSlice, 5.25)
	fmt.Println("floatSlice:", floatSlice, len(floatSlice), cap(floatSlice))

	structSlice := make([]MyStruct, 0, 10)
	structSlice = append(structSlice, MyStruct{Number: 50, String: "str"})
	fmt.Println("structSlice:", structSlice, len(structSlice), cap(structSlice))
}
//2
//— Создать несколько разных слайсов разных типов
//— Каждый слайс нужно содавать по-своему: какой-то через создание на месте, какой-то через make без выделения памяти, какой-то через make с выеделением памяти, какой-то через make с автоматическим созданием элементов со значениями по-умолчанию
//— Проделать все изученные над слайсов операции (добавление нового элемента, получение отдельного элемента, вывод на экран, изменение отдельного элемента, слайсы + циклы)