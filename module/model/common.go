package model

type RegistData struct {
	Pwd     string
	Account string
	Code    string
}

type HttpResponseJson struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
