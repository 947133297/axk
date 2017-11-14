package http

import (
	"net/http"

	"../../util"
	"../model"
)

// HandleHTTP ：处理htpp 连接
func HandleHTTP(httpHost string) {
	http.Handle("/", http.FileServer(http.Dir("log")))
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/register", register)
	util.Println("http server runing on " + httpHost)
	util.Fatalln(http.ListenAndServe(httpHost, nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test exec"))
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		key, code := util.GetCode()
		(&model.HttpResponseJson{
			Code: 0,
			Msg:  key + " - " + code,
		}).ResponseJson(w)
		return
	}

	//POST or other
	r.ParseForm()
	account := r.PostFormValue("account")
	pwd := r.PostFormValue("pwd")
	key := r.PostFormValue("key")
	code := r.PostFormValue("code")
	pass := util.VerifyCode(key, code)
	var resp *model.HttpResponseJson
	if pass {
		resp = &model.HttpResponseJson{
			Code: 0,
			Msg:  account + " - " + pwd + " success",
		}
	} else {
		resp = &model.HttpResponseJson{
			Code: 1,
			Msg:  "code err",
		}
	}
	resp.ResponseJson(w)
}
