package main

import (
	"fmt"
	"time"
)

func main() {
	//01定时2s,阻塞2s，2s后产生一个事件，往channel写内容
	<-time.After(5 / 2 * time.Second)
	fmt.Println("01时间到")
	<-time.After(3 * time.Second)
	fmt.Println("3s时间到")
	//02
	time.Sleep(time.Second * 2)
	fmt.Println("02时间到")
	//03
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("03时间到")
}
