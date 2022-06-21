package main

import (
	"log"
	"os"
)

func init() {
	// 获取日志文件句柄
	// 以 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	logFile, err := os.OpenFile(`/Users/jackmok/go/src/awesomeProject/2022/learnlog/日志文件.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 设置存储位置
	log.SetOutput(logFile)
}

func main() {
	log.Println(`今天我学会了输出日志到日志文件2`)
	log.Println(`然后我就试了一下1`)
}
