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
	UserList  []*User
	PageTitle string
}

type UserMainPageData struct {
	*HttpResponseJson
	PageTitle string
	Projects  []*Project
}

func GetHttpResponseJson(code int, msg string) *HttpResponseJson {
	return &HttpResponseJson{
		Msg:  msg,
		Code: code,
	}
}

type User struct {
	Id      int
	Role    int
	Account string
}

type Project struct {
	Id   int
	Uid  int
	Name string
}
