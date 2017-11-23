package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var logger *log.Logger

// InitLog 初始化日志
func InitLog(dirName, logName string) (err error) {
	dir, _ := os.Getwd()
	os.Mkdir(dir+"/"+dirName, os.ModePerm)
	logName = dirName + "/" + logName
	var file *os.File
	if _, err = os.Stat(logName); os.IsNotExist(err) {
		file, err = os.Create(logName)
	} else {
		file, err = os.OpenFile(logName, os.O_APPEND, 0644)
	}
	if err == nil {
		logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	}
	return
}

func Println(v ...interface{}) {
	logger.Println(v)
}

func Fatalln(v ...interface{}) {
	logger.Fatalln(v)
}

func HandleErr(err error, ifExit bool) {
	if err != nil {
		if ifExit {
			Fatalln(err.Error())
		}
		Println(err.Error())
	}
}

func EncodeFileName(name, ext string) string {
	h := md5.New()
	io.WriteString(h, name+time.Now().Format("2006-01-02 15:04:05"))
	str := fmt.Sprintf("%x", h.Sum(nil))
	return str + "." + ext
}
