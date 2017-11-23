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

type ProjectPageData struct {
	*HttpResponseJson
	*Project
	SectionList []*Section
	// 这里添加其他项目主页的数据
}

type AddSectionData struct {
	Name  string
	Graph string
	// TODO 这里添加其他检测区域数据（施工图除外）
}

// for inner dispatch
type Section struct {
	Id    string
	Pid   string
	Uid   int
	Graph string
	Name  string
}
