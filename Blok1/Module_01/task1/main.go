package main

import (
	"fmt"
)

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type HealthCheck struct {
	ServiceID int
	status    string
}

func GenerateCheck() []HealthCheck {
	var checks []HealthCheck
	for i := 0; i <= 5; i++ {
		status := FailStatus
		if i%2 == 0 {
			status = PassStatus
		}
		checks = append(checks, HealthCheck{
			ServiceID: i,
			status:    status,
		})
	}
	return checks
}

func main() {
	fmt.Println("Тут будет выведен идентификатор")

	checks := GenerateCheck()

	for _, check := range checks {
		if check.status == PassStatus {
			fmt.Println(check.ServiceID)
		}
	}
}
