package main

import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a1 int
	b1 bool
)
var c1, d1 int = 1, 2
var e1, f1 = 123, "hello"

func main() {
	var a string = "hoo.com"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println(a)

	// 没有初始化就为零值
	var d int
	fmt.Println(d)

	// bool 零值为 false
	var e bool
	fmt.Println(e)
	var i int
	var f float64
	var b1 bool
	var s string = "what is this"
	var d1 = true
	s1 := "test1"
	g, h := 123, "hello"
	println(x, y, a1, b1, c1, d1, e1, f1, g, h, "~~~\n")
	fmt.Println(s1)
	fmt.Println(d1)
	fmt.Printf("%v %v %v %q\n", i, f, b1, s)
}
