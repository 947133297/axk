package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const httpHost, tcpHost = ":12306", ":12307"

var logger *log.Logger

func main() {
	if err := initLog("log", "log.txt"); err != nil {
		panic(err)
	}

	go func() {
		handleTcpConn() // tcp srv
	}()

	http.Handle("/", http.FileServer(http.Dir("log")))
	http.HandleFunc("/test", testHandler)

	// http srv
	logger.Println("http server runing on " + httpHost)
	logger.Fatalln(http.ListenAndServe(httpHost, nil))
}
func handleTcpConn() {
	ln, err := net.Listen("tcp", tcpHost)
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Println("tcp server runing on " + tcpHost)
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Println(err)
		}
		go handleConn(conn)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test exec"))
}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	ip := conn.RemoteAddr().String()
	logger.Println("connected in " + ip)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logger.Println("error occur in read " + ip + " " + err.Error())
			return
		}
		conn.Write([]byte(fmt.Sprintf("%d", n)))
	}
}

// 初始化日志
func initLog(dirName, logName string) (err error) {
	dir, _ := os.Getwd()
	os.Mkdir(dir+"/"+dirName, os.ModePerm)
	logName = dirName + "/" + logName
	var file *os.File
	if _, err = os.Stat(logName); os.IsNotExist(err) {
		fmt.Println("create")
		file, err = os.Create(logName)
	} else {
		file, err = os.OpenFile(logName, os.O_APPEND, 0644)
	}
	if err == nil {
		logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	}
	return
}
