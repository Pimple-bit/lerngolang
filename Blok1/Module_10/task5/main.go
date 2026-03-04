package main

import "fmt"

type ParkingPlace struct {
	Address string
	Price   int
}

func main() {
	// parkingPlaces := map[string]int{
	// 	"A1":  400,
	// 	"B2":  1000,
	// 	"D10": 2100,
	// }

	// for k, v := range parkingPlaces {
	// 	if v < 500 {
	// 		fmt.Println("Address:", k, "Price:", v)
	// 	}
	// }

	// for k, v := range parkingPlaces {
	// 	if v > 900 {
	// 		parkingPlaces[k] = int(float64(v) * 0.9)
	// 	}
	// }

	// fmt.Println("parking places after discount:", parkingPlaces)

	parkingPlaces := []ParkingPlace{
		{
			Address: "A1",
			Price:   400,
		},
		{
			Address: "B2",
			Price:   300,
		},
		{
			Address: "C5",
			Price:   1100,
		},
	}

	for _, v := range parkingPlaces {
		if v.Price < 500 {
			fmt.Println("parking place with price under 500:", v.Address)
		}
	}

	for i, v := range parkingPlaces {
		if v.Price > 900 {
			parkingPlaces[i].Price = int(float64(v.Price) * 0.9)
		}
	}

	fmt.Println("parking places:", parkingPlaces)
}
