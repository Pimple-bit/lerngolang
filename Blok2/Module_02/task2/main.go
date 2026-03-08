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