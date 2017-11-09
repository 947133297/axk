package main

import "net/http"
import "fmt"

const host = ":12306"

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/", testHandler)
	fmt.Println("server runing on " + host)
	http.ListenAndServe(host, nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test exec"))
}
