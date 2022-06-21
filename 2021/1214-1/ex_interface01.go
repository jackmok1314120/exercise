package main

import "fmt"

type Animal interface {
	Eat()
	Run()
}

type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating\n", d.Name)
}

func (d *Dog) Run() {
	fmt.Printf("%s is running\n", d.Name)
}

func ShowEat(animal Animal) {
	animal.Eat()
}

func ShowRun(animal Animal) {
	animal.Run()
}

func main() {
	dog := Dog{Name: "Kenny"}
	ShowEat(&dog)
	ShowRun(&dog)
	/*var animal Animal
	fmt.Println(animal) //底层型态为nil
	animal = &Dog{Name:"Kenny"}
	fmt.Println(animal) */
}
