package http

import (
	"fmt"
	"math/rand"
	"net/http"
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
	app.HttpServer.ServerFile("/img/*filepath", wd+"/static/vk")
	app.HttpServer.SetEnabledListDir(false)

	// 配置路由
	app.HttpServer.POST("/register", register)
	app.HttpServer.GET("/vk", getVK)

	// 开始运行
	util.Println("http server runing on :" + strconv.Itoa(httpHost))
	panic(app.StartServer(httpHost))
}

func register(ctx dotweb.Context) error {
	data := new(model.RegistData)
	if err := ctx.Bind(data); err != nil {
		return err
	}
	k := ctx.Session().Get("vk")
	var key string
	if k != nil {
		key = k.(string)
	} else {
		util.Println("session read failed, get nil", "\r\n")
	}
	pass := util.VerifyCode(key, data.Code)
	resp := new(model.HttpResponseJson)
	if !pass {
		resp.Code = 1
		resp.Msg = "verify code err"
	} else {
		resp.Code = 0
		resp.Msg = "ok"
		// TODO 执行入库代码
	}
	ctx.WriteJson(resp)
	return nil
}

func getVK(ctx dotweb.Context) error {
	rc := strconv.Itoa(rand.Intn(9) + 1)
	if err := ctx.Session().Set("vk", rc); err != nil {
		util.Println("session set error => ", err, "\r\n")
	}
	url := fmt.Sprintf("/img/%s.jpg", rc)
	return ctx.Redirect(http.StatusMovedPermanently, url)
}
