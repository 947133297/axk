package http

import (
	"fmt"
	"os"
	"strconv"

	"../../util"
	"../model"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/session"
)

// HandleHTTP ：处理htpp 连接
func HandleHTTP(httpHost int) {
	app := dotweb.New()

	// 会话配置
	app.HttpServer.SetEnabledSession(true)
	app.HttpServer.SetSessionConfig(session.NewDefaultRuntimeConfig())

	// 配置静态文件
	wd, _ := os.Getwd()
	app.HttpServer.ServerFile("/web/*filepath", wd+"/static/web")
	app.HttpServer.SetEnabledListDir(false)

	// 配置路由
	app.HttpServer.Any("/register", register)

	// 开始运行
	util.Println("http server runing on :" + strconv.Itoa(httpHost))
	panic(app.StartServer(httpHost))
}

func register(ctx dotweb.Context) error {
	r := ctx.Request()
	if r.Method == "GET" {
		key, code := util.GetCode()
		ctx.WriteJson(&model.HttpResponseJson{
			Code: 0,
			Msg:  key + " - " + code,
		})
		return nil
	}

	//POST or other
	// r.ParseForm()
	// account := r.PostFormValue("account")
	// pwd := r.PostFormValue("pwd")
	// key := r.PostFormValue("key")
	// code := r.PostFormValue("code")

	// pass := util.VerifyCode(key, code)
	data := new(model.RegistData)
	if err := ctx.Bind(data); err != nil {
		return err
	}
	ctx.WriteString(fmt.Sprint(data))
	return nil
}
