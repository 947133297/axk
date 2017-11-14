package main

import (
	"./module/http"
	"./module/tcp"
	"./util"
)

const httpHost, tcpHost = ":12306", ":12307"

func main() {
	if err := util.InitLog("log", "log.txt"); err != nil {
		panic(err)
	}

	go func() {
		tcp.HandleTCP(tcpHost)
	}()
	http.HandleHTTP(httpHost)
}
