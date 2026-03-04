package x

import "fmt"

var privateXGlobal = "private"

func privateX() {
	fmt.Println("i'm X, and i'm private function!")
}

type privateXStruct struct{}

func (s privateXStruct) privateXMethod() {
	fmt.Println("i'm x, and i'm private method")
}
