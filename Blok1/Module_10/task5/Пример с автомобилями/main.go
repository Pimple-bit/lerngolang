package main

import "fmt"

type Auto struct {
	Number  string
	Invalid bool
}

type ParkingPlace struct {
	Price   int
	Invalid bool

	Auto *Auto
}

func main() {
	auto1 := Auto{
		Number:  "A111AA",
		Invalid: true,
	}

	auto2 := Auto{
		Number:  "B222BB",
		Invalid: false,
	}

	parkingPlaces := map[string]ParkingPlace{
		"A1": {
			Price:   400,
			Invalid: true,
			Auto:    &auto1,
		},
		"B2": {
			Price:   1000,
			Invalid: false,
			Auto:    nil,
		},
		"D10": {
			Price:   500,
			Invalid: true,
			Auto:    &auto2,
		},
	}

	parkingPlace, ok := parkingPlaces["D10"]
	if !ok {
		fmt.Println("такого парковочного места нет!")
		return
	}

	if parkingPlace.Auto == nil {
		fmt.Println("на парковочном месте нет автомобиля")
		return
	}

	if parkingPlace.Invalid {
		fmt.Println("парковочное место предназначено для инвалидов")
		if parkingPlace.Auto.Invalid {
			fmt.Println("всё хорошо, на инвалидном месте стоит инвалид")
		} else {
			fmt.Println("на инвалидном месте стоит НЕ инвалид! эвакуировать!")
		}
	} else {
		fmt.Println("парковочное место НЕ предназначено для инвалидов")
	}
}
