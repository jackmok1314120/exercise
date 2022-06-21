package main

import (
	"log"
)

type User struct {
	Name string
	Age  int
}

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
)

func main() {
	u := User{
		Name: "dj",
		Age:  18,
	}

	//log.Printf("%s login, age:%d", u.Name, u.Age)
	//log.Panicf("Oh, system error when %s login", u.Name)
	//log.Fatalf("Danger! hacker %s login", u.Name)

	//log.SetPrefix("Login: ") //调用前缀
	//log.Printf("%s login, age:%d", u.Name, u.Age)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds) //
	log.Printf("%s login, age:%d", u.Name, u.Age)
}
