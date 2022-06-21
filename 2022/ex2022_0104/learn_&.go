package main

import "fmt"

func main() {
	var a int = 10 /* 声明实际变量 */
	var b *int     /* 声明指针变量 */
	b = &a         /* 指针变量的存储地址 */

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	fmt.Printf("变量的地址: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("b变量储存的指针地址: %x\n", b)
	/* 使用指针访问值 */
	fmt.Printf("*b的值: %d\n", *b)
}
