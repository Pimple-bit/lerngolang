package x

import "fmt"

var PublicXGlobal = 100

func PublicX() {
	fmt.Println("i'm X, and i'm public function")
}

type PublicXStruct struct{}

func (s PublicXStruct) PublicXMethod() {
	fmt.Println("i'm x, and i'm public method")
}

func (s PublicXStruct) PublicXMethod2() {
	fmt.Println("private var:", privateXGlobal)
	privateX()

	privateStruct := privateXStruct{}
	privateStruct.privateXMethod()
}
