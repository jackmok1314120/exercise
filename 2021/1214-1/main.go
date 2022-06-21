package main

import (
	"fmt"
) //.fmt
func main() {
	/*fmt.Println("hi" + "jack")
	var stockcode=123
	var enddate="2020-12-31"
	var url="Code=%d&endDate=%s"
	var target_url=fmt.Sprintf(url,stockcode,enddate)
	fmt.Println(target_url)*/
	//fmt.Println(math.Exp2(10))  // 1024
	a, b, c := "", 0, false
	/*fmt.Print("hi")
	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	fmt.Println("a", "b", 1, 2, 3, "c", "d")
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)*/
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, c)
	//if err := percent(30, 70, 90, 160); err != nil {
	//	fmt.Println(err)
}

//}
/*func percent(i ...int) error {
	for _, n := range i {
		if n > 100 {
			return fmt.Errorf("数值 %d 超出范围（100）", n)
		}
		fmt.Print(n, "%\n")
	}
	return nil
}*/
