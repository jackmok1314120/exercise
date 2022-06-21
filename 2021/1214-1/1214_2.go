package main

import (
	"fmt"
)

type Phone interface {
	call() string
}

type Android struct {
	brand string
}

type IPhone struct {
	version string
}

func (android Android) call() string {
	return "I am Android " + android.brand
}

func (iPhone IPhone) call() string {
	return "I am iPhone " + iPhone.version
}

func printCall(p Phone) {
	fmt.Println(p.call() + ", I can call you!")
}

func main() {
	var vivo = Android{brand: "Vivo"}
	var hw = Android{"HuaWei"}

	i7 := IPhone{"7 Plus"}
	ix := IPhone{"X"}

	printCall(vivo)
	printCall(hw)
	printCall(i7)
	printCall(ix)
}
