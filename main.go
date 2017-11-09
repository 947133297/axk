package main

import (
	"fmt"
	"net"
	"net/http"
)

const httpHost, tcpHost = ":12306", ":12307"

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))
	http.HandleFunc("/", testHandler)

	// http srv
	fmt.Println("http server runing on " + httpHost)
	go http.ListenAndServe(httpHost, nil)

	// tcp srv
	ln, err := net.Listen("tcp", tcpHost)
	if err != nil {
		panic(err)
	}
	fmt.Println("tcp server runing on " + tcpHost)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConn(conn)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test exec"))
}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	fmt.Println("connected in .. " + conn.RemoteAddr().String())
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		conn.Write([]byte(fmt.Sprintf("%d", n)))
	}
}
