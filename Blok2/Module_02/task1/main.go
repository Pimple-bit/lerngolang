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