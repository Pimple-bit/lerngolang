package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printHelp() {
	fmt.Println("добавить {ингридиент} {количество} -- команда для добавления ингридиента")
	fmt.Println("удалить {ингридиент} {количество} -- команда для удаления ингридиента")
	fmt.Println("получить {ингридиент} -- команда для получения ингридиента")
	fmt.Println("help -- команда для получения списка доступных команд")
	fmt.Println("выйти -- команда для выхода из программы")
	fmt.Println("")
}

func main() {
	ingridients := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")

		ok := scanner.Scan()
		if !ok {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)
		if len(fields) == 0 {
			continue
		}

		cmd := fields[0]

		if cmd == "выйти" {
			fmt.Println("Вы завершили программу!")
			return
		}

		if cmd == "help" {
			printHelp()
			continue
		}

		if cmd == "добавить" {
			if len(fields) != 3 {
				fmt.Println("Команда имеет неверный формат!")
				continue
			}

			amount, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("Неверный формат количества ингридиента!")
				continue
			}

			if amount <= 0 {
				fmt.Println("Число, на которое меняется ингридиент, должно быть положительным!")
			}

			ingridients[fields[1]] += amount
			continue
		}

		if cmd == "удалить" {
			if len(fields) != 3 {
				fmt.Println("Команда имеет неверный формат!")
				continue
			}

			amount, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("Неверный формат количества ингридиента!")
				continue
			}

			ingridientAmount, ok := ingridients[fields[1]]
			if !ok {
				fmt.Println("Попытка удалить несуществующий ингридиент!")
				continue
			}

			if ingridientAmount-amount < 0 {
				fmt.Println("Попытка удалить больше ингридиента, чем было добавлено!")
				continue
			}

			if amount <= 0 {
				fmt.Println("Число, на которое меняется ингридиент, должно быть положительным!")
			}

			ingridients[fields[1]] -= amount
			if ingridients[fields[1]] == 0 {
				delete(ingridients, fields[1])
			}
			continue
		}

		if cmd == "получить" {
			if len(fields) != 2 {
				fmt.Println("Команда имеет неверный формат!")
				continue
			}

			ingridientAmount, ok := ingridients[fields[1]]
			if !ok {
				fmt.Println("Данный ингридиент не найден!")
				continue
			}

			fmt.Println("Ингридиент:", fields[1], "Количество:", ingridientAmount)
			continue
		}

		fmt.Println("Передана неизвестная команда!")
	}
}
