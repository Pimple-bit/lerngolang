package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type BankAccount struct {
	balance float64
}

func randomError() error {
	if rand.Float64() < 0.3 {
		return errors.New("техническая ошибка выполнения операции")
	}
	return nil
}

func (a *BankAccount) Withdraw(amount float64) error {

	if err := randomError(); err != nil {
		return fmt.Errorf("ошибка снятия наличных: %w", err)
	}

	if amount > a.balance {
		return errors.New("недостаточно средств")
	}

	a.balance -= amount
	fmt.Println("Выдано наличных:", amount)

	return nil
}

func (a *BankAccount) PayOnline(amount float64) error {

	if err := randomError(); err != nil {
		return fmt.Errorf("ошибка онлайн оплаты: %w", err)
	}

	if amount > a.balance {
		return errors.New("недостаточно средств")
	}

	a.balance -= amount
	fmt.Println("Онлайн оплата:", amount)

	return nil
}

func (a *BankAccount) ShowBalance() error {

	if err := randomError(); err != nil {
		return fmt.Errorf("ошибка получения баланса: %w", err)
	}

	fmt.Println("Текущий баланс:", a.balance)
	return nil
}

func main() {

	rand.Seed(time.Now().UnixNano())

	account := BankAccount{
		balance: 1000,
	}

	for {

		operation := rand.Intn(3)
		amount := float64(rand.Intn(500))

		var err error

		switch operation {

		case 0:
			err = account.Withdraw(amount)

		case 1:
			err = account.PayOnline(amount)

		case 2:
			err = account.ShowBalance()
		}

		if err != nil {
			fmt.Println("Ошибка:", err)
		}

		time.Sleep(time.Second)
	}
}
//— Необходимо реализовать работу банковского счёта, который:
//1. Умеет выдавать наличные
//2. Умеет показывать текущий баланс
//3. Позволяет проводить онлайн-оплаты
//— Помимо ошибки "Недостаточно средств", при проведении каждой операции: выдача наличных, показ текущего баланса, онлайн-оплата, с вероятностью 30% происходит ошибка проведения операции
//— Нужно в бесконечном цикле реализовать симуляцию проведения операций с банковским счётом
//— Необходимо отлавливать возникающие ошибки, и выводить их на экран, с подробным описанием самой ошибки и причины её возникновения