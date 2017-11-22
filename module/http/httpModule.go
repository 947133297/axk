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
	app.HttpServer.ServerFile("/static/*filepath", wd+"/static")
	app.HttpServer.SetEnabledListDir(false)

	// 配置路由
	app.HttpServer.POST("/register", register)
	app.HttpServer.POST("/login", login)
	app.HttpServer.GET("/vk", getVK)
	app.HttpServer.GET("/getMgrData", getMgrData)
	app.HttpServer.GET("/getUserData", getUserData)

	// 开始运行
	util.Println("http server runing on :" + strconv.Itoa(httpHost))
	panic(app.StartServer(httpHost))
}
func login(ctx dotweb.Context) error {
	loginData := new(model.LoginData)
	var err error
	if err = ctx.Bind(loginData); err != nil {
		return err
	}
	user, err := util.UserLogin(loginData)
	if user != nil {
		if err = ctx.Session().Set("user", user); err != nil {
			util.Println("session set user error => ", err, "\r\n")
		} else {
			// 返回角色id，让前端选择路由
			ctx.WriteJson(model.GetHttpResponseJson(0, strconv.Itoa(user.Role)))
			return nil
		}
	}
	ctx.WriteJson(model.GetHttpResponseJson(1, err.Error()))
	return nil
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
	if !pass {
		ctx.WriteJson(model.GetHttpResponseJson(1, "verify code err"))
	} else {
		err := util.RegisteUser(data)
		if err == nil {
			ctx.WriteJson(model.GetHttpResponseJson(0, "ok"))
		} else {
			ctx.WriteJson(model.GetHttpResponseJson(1, err.Error()))
		}
	}
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

// 获取管理员主页数据 for ajax
func getMgrData(ctx dotweb.Context) error {
	user := fetchSessionData(ctx)
	// TODO 设置数据帧结构，返回管理员数据
	data := new(model.MgrMainPageData)
	if user == nil {
		data.HttpResponseJson = model.GetHttpResponseJson(2, "to login")
		ctx.WriteJson(data)
	} else if user.Role != 1 {
		data.HttpResponseJson = model.GetHttpResponseJson(3, "no auth")
		ctx.WriteJson(data)
	} else {
		data.HttpResponseJson = model.GetHttpResponseJson(0, "ok")
		data.PageTitle = "管理员页面 : " + user.Account
		data.UserList = util.GetUserList()
		ctx.WriteJson(data)
	}
	return nil
}

// 从session中获取用户数据
func fetchSessionData(ctx dotweb.Context) (user *model.User) {
	u := ctx.Session().Get("user")
	if u == nil {
		return
	}
	user = u.(*model.User)
	return
}

func getUserData(ctx dotweb.Context) error {
	user := fetchSessionData(ctx)
	if user.Role == 1 {
		// 管理员查看普通用户
		id, err := strconv.Atoi(ctx.QueryString("u"))
		if err != nil {
			ctx.WriteJson(model.GetHttpResponseJson(1, "uid err"))
			return nil
		}
		user, err = util.GetUser(id)
		if err != nil {
			util.Println(err.Error())
			user = nil
		}
	}
	if user == nil {
		ctx.WriteJson(model.GetHttpResponseJson(1, "user nil"))
		return nil
	}
	// 获取user数据返回
	data := new(model.UserMainPageData)
	data.HttpResponseJson = model.GetHttpResponseJson(0, "ok")
	data.Projects = util.GetAllProject(user.Id)
	data.PageTitle = "普通用户 ： " + user.Account
	ctx.WriteJson(data)
	return nil
}
