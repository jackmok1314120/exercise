package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	s := []string{"sting", "abc"}
	sa := []string{"def", "ghi"}
	//注意下面这两个区别
	fmt.Println(append(x, 4, 5, 6)) //slice = append(slice, elem1, elem2)
	fmt.Println(append(x, y...))    //slice = append(slice, anotherSlice...)
	fmt.Println(append(s, sa...))
	/*result:
	  [1 2 3 4 5 6]
	  [1 2 3 4 5 6]
	  [sting abc def ghi]
	*/
}
