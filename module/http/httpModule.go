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

	filePath = wd + "/static/upload"

	// 配置路由
	app.HttpServer.POST("/register", register)
	app.HttpServer.POST("/login", login)
	app.HttpServer.GET("/vk", getVK)
	app.HttpServer.GET("/getMgrData", getMgrData)
	app.HttpServer.GET("/getUserData", getUserData)
	app.HttpServer.GET("/addProject", addProject)
	app.HttpServer.GET("/getProjectData", getProjectData)
	app.HttpServer.POST("/addWatchSection", addWatchSection)
	app.HttpServer.POST("/upload", upload)

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
	user := fetchActualUser(ctx)
	if user == nil {
		return nil
	}
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
func fetchActualUser(ctx dotweb.Context) (user *model.User) {
	u := ctx.Session().Get("user")
	if u == nil {
		ctx.WriteJson(model.GetHttpResponseJson(2, "to login"))
		return
	}
	user = u.(*model.User)
	id, err := strconv.Atoi(ctx.QueryString("u"))
	if user.Role == 1 && err == nil && id > 0 {
		user, err = util.GetUser(id)
		if err != nil {
			util.Println(err.Error())
		}
	}
	if user == nil {
		ctx.WriteJson(model.GetHttpResponseJson(1, "user nil"))
	}
	return
}

func getUserData(ctx dotweb.Context) error {
	user := fetchActualUser(ctx)
	if user == nil {
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

func addProject(ctx dotweb.Context) error {
	user := fetchActualUser(ctx)
	if user == nil {
		return nil
	}
	err := util.AddProject(user.Id, ctx.QueryString("pjname"))
	if err != nil {
		util.Println(err.Error())
		ctx.WriteJson(model.GetHttpResponseJson(1, err.Error()))
	} else {
		ctx.WriteJson(model.GetHttpResponseJson(0, "ok"))
	}
	return nil
}

func getProjectData(ctx dotweb.Context) error {
	user := fetchActualUser(ctx)
	if user == nil {
		return nil
	}
	pid := ctx.QueryString("p")
	project := util.GetProjectMsg(user.Id, pid)
	if project == nil {
		ctx.WriteJson(model.GetHttpResponseJson(1, "project nil"))
		return nil
	}

	data := new(model.ProjectPageData)
	data.Project = project
	data.HttpResponseJson = model.GetHttpResponseJson(0, "ok")
	data.SectionList = util.GetAllSection(pid, user.Id)
	ctx.WriteJson(data)
	return nil
}

var filePath string

func upload(ctx dotweb.Context) error {
	user := fetchActualUser(ctx)
	if user == nil {
		return nil
	}
	file, err := ctx.Request().FormFile("file")
	if err != nil {
		util.Println(err.Error())
		return err
	}
	fileName := util.EncodeFileName(file.FileName(), file.GetFileExt())
	_, err = file.SaveFile(filePath + "/" + fileName)
	if err != nil {
		util.Println(err.Error())
		return err
	}
	ctx.WriteJson(model.GetHttpResponseJson(0, fileName))
	return nil
}

// needs p u params
func addWatchSection(ctx dotweb.Context) error {
	user := fetchActualUser(ctx)
	if user == nil {
		return nil
	}

	section := new(model.AddSectionData)
	var err error

	// 获取基本数据
	if err = ctx.Bind(section); err != nil {
		util.Println(err.Error())
		return err
	}

	// 入库
	sd := &model.Section{
		Pid:   ctx.QueryString("p"),
		Uid:   user.Id,
		Name:  section.Name,
		Graph: section.Graph,
	}
	err = util.AddSection(sd)
	if err != nil {
		util.Println(err.Error())
		ctx.WriteJson(model.GetHttpResponseJson(1, err.Error()))
		return nil
	}
	// 响应
	ctx.WriteJson(model.GetHttpResponseJson(0, "ok"))
	return nil
}
