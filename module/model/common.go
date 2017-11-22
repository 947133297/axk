package model

type RegistData struct {
	Pwd     string
	Account string
	Code    string
}
type LoginData struct {
	Pwd     string
	Account string
}

type HttpResponseJson struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
type MgrMainPageData struct {
	*HttpResponseJson
	UserList  []string
	PageTitle string
}

func GetHttpResponseJson(code int, msg string) *HttpResponseJson {
	return &HttpResponseJson{
		Msg:  msg,
		Code: code,
	}
}

type User struct {
	Role    int
	Account string
}
