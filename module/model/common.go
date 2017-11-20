package model

type RegistData struct {
	Pwd     string
	Key     string
	Code    string
	Account string
}

type HttpResponseJson struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
