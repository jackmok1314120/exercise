package main

import (
	"fmt"
	"runtime/debug"
)

func fun() {
	fmt.Println("fun begin")

	defer func() {
		//捕获panic
		if err := recover(); err != nil {
			debug.PrintStack()
			//获取堆栈信息的字符串
			fmt.Println("xxx", string(debug.Stack()))
		}
	}()

	var p *int
	//产生异常
	*p = 0
	fmt.Println("fun end")
	//这里不执行
	for {
	}
}

func main() {
	fmt.Println("main begin")
	fun()
	//因为panic被recover捕获，所以下面继续执行
	fmt.Println("main end")

	for {
	}
}
