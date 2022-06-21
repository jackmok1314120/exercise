package main

import (
	"fmt"
	"strings"
)

var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

func main() {
	str := "hi \nhoo\n.com"
	fmt.Println("-------- 原字符串 ----------")
	fmt.Println(str)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println("-------- 去除空格与换行后 ----------")
	fmt.Println(str)
	//系統自動識別數字類型
	var a = 1.5
	var b = 2
	fmt.Println(a, b)
	//var available bool  // 一般声明
	//valid := false      // 简短声明
	//available = true    // 赋值操作
}
