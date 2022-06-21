/*
* Copyright(C),2019-2020, email: 952651117@qq.com
* Author:  dao
* Version: 1.0.0
* Date:    2021/6/11 13:54
* Description:
*
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var logFile io.Writer

/**
 * @Description:判断所给路径是否为文件
 * @param path
 * @return os.FileInfo
 * @return bool
 */

func IsFile(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && !f.IsDir()
}

/**
 * @Description:判断路径是否存在
 * @param path
 * @return os.FileInfo
 * @return bool
 */

func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

/**
 * @Description: 写日志
 * @param msg
 * @param title
 */

func Logwrite() {
	//设置时间变量
	now := time.Now()
	//前提：创建一个 log 文件夹到根目录（这样最好了）。如果文件目录不存在，也可以使用下面代码去创建文件目录
	folderName := "log/" + now.Format("20060102")
	folderTemp := ""
	for _, v := range strings.Split(folderName, "/") {
		//判断文件，如果不存在，那么就创建
		folderTemp += v
		_, IsExistsPath := IsExists(folderTemp)
		if !IsExistsPath {
			//创建文件夹代码
			os.Mkdir(folderTemp, os.ModePerm)
		}
		folderTemp += "/"
	}

	//设置日志文件，做好日志文件管理：这里是用时间做文件，每个小时生成不同文件
	filefullname := "./" + folderName + "/" + now.Format("20060102") + now.Format("_15") + ".log"
	//判断文件是否存在
	_, b := IsFile(filefullname)
	if b {
		//打开文件，
		logFile, _ = os.OpenFile(filefullname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	} else {
		//新建文件
		logFile, _ = os.Create(filefullname)
	}

}

var Tag = []string{
	"api",
}

func InitRouter(r *gin.Engine, g []string) {
	for i, _ := range g {
		matchRouter(g[i], r)
		//fmt.Println("打印g[i]:", g[i])
	}

}
func matchRouter(tag string, r *gin.Engine) {
	path := fmt.Sprintf("%s/%s", strings.ToLower(tag), strings.ToLower(""))
	switch tag {
	case "api":
		LoadCallBackRouter(r.Group(path))
		break
	}
}

func LoadCallBackRouter(group *gin.RouterGroup) {
	group.POST("/callback", callback)
}

func callback(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")
	ctx.Header("content-type", "application/json")
	//ctx.JSON(http.StatusOK, map[string]interface{}{})
	data := map[string]interface{}{}

	ctx.BindJSON(&data)
	//log.Println(data)
	//fmt.Println("收到的data:", data)
	//log.Fatal("%v", &data)
	ctx.JSON(http.StatusOK, gin.H{
		"massage": http.StatusOK,
		"data":    data,
	})
	return
}

func main() {
	r := gin.Default()
	InitRouter(r, Tag)
	r.Run(":9081")
	//fmt.Println()
}
