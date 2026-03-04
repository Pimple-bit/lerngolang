package main

import "github.com/k0kubun/pp"

type Dog struct {
	Name       string
	Estimation int
	Respectful bool
}

func main() {
	dogs := []Dog{
		{
			Name:       "Белла",
			Estimation: 7,
			Respectful: true,
		},
		{
			Name:       "Джесси",
			Estimation: 6,
			Respectful: true,
		},
		{
			Name:       "Умка",
			Estimation: 8,
			Respectful: false,
		},
		{
			Name:       "Роза",
			Estimation: 5,
			Respectful: false,
		},
		{
			Name:       "Белка",
			Estimation: 5,
			Respectful: true,
		},
	}

	pp.Println("Собаки до:", dogs)

	for i, v := range dogs {
		if v.Respectful {
			dogs[i].Estimation++
		}
	}

	pp.Println("Собаки после:", dogs)
}
