package main

import (
	"errors"
	"fmt"
)

func checkArgs(a, b float64) error {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return errors.New("аргументы должны быть в диапазоне [-1000, 1000]")
	}
	return nil
}

func add(a, b float64) (float64, error) {
	if err := checkArgs(a, b); err != nil {
		return 0, err
	}
	return a + b, nil
}

func sub(a, b float64) (float64, error) {
	if err := checkArgs(a, b); err != nil {
		return 0, err
	}
	return a - b, nil
}

func mul(a, b float64) (float64, error) {
	if err := checkArgs(a, b); err != nil {
		return 0, err
	}
	return a * b, nil
}

func div(a, b float64) (float64, error) {
	if err := checkArgs(a, b); err != nil {
		return 0, err
	}

	if b == 0 {
		return 0, errors.New("деление на 0")
	}

	return a / b, nil
}

func printResult(op string, a, b float64, result float64, err error) {
	if err != nil {
		fmt.Printf("Ошибка операции %s с аргументами %.2f и %.2f: %v\n", op, a, b, err)
	} else {
		fmt.Printf("Результат %s (%.2f, %.2f) = %.2f\n", op, a, b, result)
	}
}

func main() {

	r, err := add(10, 5)
	printResult("add", 10, 5, r, err)

	r, err = sub(2000, 5)
	printResult("sub", 2000, 5, r, err)

	r, err = mul(10, -2000)
	printResult("mul", 10, -2000, r, err)

	r, err = div(10, 0)
	printResult("div", 10, 0, r, err)

	r, err = div(20, 4)
	printResult("div", 20, 4, r, err)
}

//— Реализуйте простую программу-калькулятор, которая поддерживает четыре операции: сложение, вычитание, умножение и деление.
//— Операции должны возвращать ошибку, если:
///1. Деление на 0.
//2. Аргументы меньше -1000 или больше 1000 (искусственное ограничение для внесения разнообразия в возможные ошибки)
//— В функции main() вызовите каждую операцию с корректными и некорректными входными данными.
//— В случае возникновения ошибки, необходимо вывести на экран операцию и значение аргументов, при которых произошла ошибка, а так же необходимо вывести саму ошибку