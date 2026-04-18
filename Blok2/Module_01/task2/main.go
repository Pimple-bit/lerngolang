package main

import "fmt"


type Notifier interface {
	Send()
}


type Email struct {
	Address string
	Message string
}

func (e Email) Send() {
	fmt.Println("Отправка Email")
	fmt.Println("Адрес:", e.Address)
	fmt.Println("Сообщение:", e.Message)
	fmt.Println("------")
}


type SMS struct {
	Phone   string
	Message string
}

func (s SMS) Send() {
	fmt.Println("Отправка SMS")
	fmt.Println("Телефон:", s.Phone)
	fmt.Println("Сообщение:", s.Message)
	fmt.Println("------")
}


type Push struct {
	DeviceID string
	Message  string
}

func (p Push) Send() {
	fmt.Println("Отправка Push")
	fmt.Println("Device:", p.DeviceID)
	fmt.Println("Сообщение:", p.Message)
	fmt.Println("------")
}


func SendNotifications(notifications []Notifier) {
	for _, n := range notifications {
		n.Send()
	}
}

func main() {
	notifications := []Notifier{
		Email{Address: "user@mail.com", Message: "Hello from Email"},
		SMS{Phone: "+123456789", Message: "Hello from SMS"},
		Push{DeviceID: "device123", Message: "Hello from Push"},
	}

	SendNotifications(notifications)
}
//— Вы разрабатываете систему отправки уведомлений через разные каналы связи:
//1. email
//2. SMS
//3. push-уведомления
//— Необходимо разработать модуль отправки уведомлений, который одним методом может принять массив уведомлений на отправку по различным каналам связи, после чего, собственно, отправить все эти уведомления
//— Вместо настоящей отправки уведомлений необходимо использовать заглушки, на манер заглушек методов оплаты для модуля проведения оплат из видео-урока
