package util

import (
	"fmt"
	"os"
	"testing"

	"../module/model"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

var pwd, account string = "123", "testtest123"

func setUp() {
	if err := InitLog("static", "log.txt"); err != nil {
		panic(err)
	}
}
func tearDown() {
	err := DelUser(account)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func TestUser(t *testing.T) {
	//注册
	err := RegisteUser(&model.RegistData{
		Pwd:     pwd,
		Account: account,
	})
	if err != nil {
		t.Error(err.Error())
	}
	// 登录
	user, err := UserLogin(&model.LoginData{
		Pwd:     pwd,
		Account: account,
	})
	if user == nil {
		t.Error("登录失败" + err.Error())
	}

	//再次注册
	err = RegisteUser(&model.RegistData{
		Pwd:     pwd,
		Account: account,
	})
	if err == nil {
		t.Error("再次注册结果错误")
	}
	fmt.Println(err.Error())
}
