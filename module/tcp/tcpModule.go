package tcp

import (
	"fmt"
	"net"

	"../../util"
)

// HandleTCP ：处理TCP 连接
func HandleTCP(tcpHost string) {
	ln, err := net.Listen("tcp", tcpHost)
	if err != nil {
		util.Fatalln(err)
	}
	util.Println("tcp server runing on " + tcpHost)
	for {
		conn, err := ln.Accept()
		if err != nil {
			util.Println(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	ip := conn.RemoteAddr().String()
	util.Println("connected in " + ip)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			util.Println("error occur in read " + ip + " " + err.Error())
			return
		}
		conn.Write([]byte(fmt.Sprintf("%d", n)))
	}
}
