package main

import ( // 调用的包
	"fmt"
	"io"
)

func main() {
	var a, b int
	for { // 循环获取输入
		if _, err := fmt.Scan(&a, &b); err != io.EOF {
			fmt.Println("a+b", a+b)
		} else {
			break
		}
	}
}
