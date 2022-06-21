package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"path"
	"time"
)

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		//rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})
	log.AddHook(lfHook)
}

//切割日志和清理过期日志
func ConfigLocalFilesystemLogger1(filePath string) {
	writer, err := rotatelogs.New(
		filePath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(filePath),           // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Second*60*3),     // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Second*60), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatal("Init log failed, err:", err)
	}
	log.SetOutput(writer)
	log.SetLevel(log.InfoLevel)
}

func main() {
	ConfigLocalFilesystemLogger1("/Users/jackmok/go/src/awesomeProject/2022/learnlog")
	for {
		log.Info(111)
		time.Sleep(500 * time.Millisecond)
	}
}
