package main

import "fmt"

type MyStruct struct {
	Number int
	String string
}

func main() {
	arrStruct := [3]MyStruct{
		{
			Number: 10,
			String: "hi",
		},
		{
			Number: 20,
			String: "hello",
		},
		{
			Number: 30,
			String: "goodbye",
		},
	}

	// for _, v := range arrStruct {
	// 	v.Number += 100
	// }

	for i := 0; i < len(arrStruct); i++ {
		arrStruct[i].Number += 100
	}

	for i, v := range arrStruct {
		fmt.Println("i:", i, " -- v:", v)
	}

	for i := range arrStruct {
		fmt.Println("I:", i)
	}

}
