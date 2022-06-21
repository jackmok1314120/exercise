package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*	var C, c int //声明变量
			C = 1        //这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）
		A:
			for C < 100 {
				C++ //C=1不能写入for这里就不能写入
				for c = 2; c < C; c++ {
					if C%c == 0 {
						goto A //若发现因子则不是素数
					}
				}
				fmt.Println(C, "是素数")
			}
	*/
	/*for i := 1; i <= 9; i++ { // i 控制行，以及计算的最大值
		for j := 1; j <= i; j++ { // j 控制每行的计算个数
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println("")
	}*/
	var score int = 0
	var fenshu string = "A"
	fmt.Printf("欢迎进入成绩查询系统\n")
MAIN:
	for true {
		var xuanzhe int = 0
		fmt.Println("1.进入成绩查询 2.退出程序")
		fmt.Printf("请输入序号选择:")
		fmt.Scanln(&xuanzhe)
		fmt.Printf("\n")
		if xuanzhe == 1 {
			goto CHA
		}
		if xuanzhe == 2 {
			os.Exit(1)
		}

	}

CHA:
	for true {
		fmt.Printf("请输入一个学生的成绩:")
		fmt.Scanln(&score)

		switch {
		case score >= 90:
			fenshu = "A"

		case score >= 80 && score < 90:
			fenshu = "B"

		case score >= 60 && score < 80:
			fenshu = "C"

		default:
			fenshu = "D"
		}

		//fmt.Println(fenshu)
		var c string = strconv.Itoa(score)
		switch {
		case fenshu == "A":
			fmt.Printf("你考了%s分,评价为%s,成绩优秀\n", c, fenshu)
		case fenshu == "B" || fenshu == "C":
			fmt.Printf("你考了%s分,评价为%s,成绩良好\n", c, fenshu)
		case fenshu == "D":
			fmt.Printf("你考了%s分,评价为%s,成绩不合格\n", c, fenshu)
		}
		fmt.Printf("\n")
		goto MAIN
	}
	//fmt.Println(score)
}
