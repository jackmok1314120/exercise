package main

import "fmt"

func main() {
	/*
	   遍历数组：依次获取数组中的数据
	   range 数组名：
	      index，value
	*/
	/*arr := [...]int{6, 2, 4, 9, 8, 3}
	//1.遍历方式一
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], "\t")
	}
	fmt.Println()

	//2.range遍历数组
	sum := 0
	for _, value := range arr {
		//fmt.Println(index, "\t",value)
		fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)*/
	i := float64(39596) / float64(100000000)

	fmt.Printf("f is %i", i)
}
