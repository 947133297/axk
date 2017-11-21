package util

import (
	"database/sql"

	"../module/model"
	_ "github.com/go-sql-driver/mysql"
)

const _DB_NAME = "axk"

var db *sql.DB

func init() {
	initDB()
}
func initDB() {
	// create db if not exists
	var err error
	db, err = sql.Open("mysql", "root:qq5566@/mysql")
	HandleErr(err, true)
	db.Exec("create database " + _DB_NAME)
	db, err = sql.Open("mysql", "root:qq5566@/"+_DB_NAME)
	HandleErr(err, true)

	//create table # USER
	// msg_tmp1 for role msg
	_, err = db.Exec(`
		create table if not exists user(
			id int(10) primary key not  null  auto_increment,
			account varchar(30),
			pwd varchar(30) not null,
			msg_tmp1 varchar(50),
			msg_tmp2 varchar(50),
			unique key (account)
		)`)
	HandleErr(err, true)
}

func UserLogin(data *model.LoginData) (user *model.User, err error) {
	var role int
	err = db.QueryRow("SELECT msg_tmp1 FROM user WHERE account=? and pwd=?", data.Account, data.Pwd).Scan(&role)
	switch {
	case err == sql.ErrNoRows:
		user = nil
	case err != nil:
		Println(err.Error())
		user = nil
	default:
		user = &model.User{
			Account: data.Account,
			Role:    role,
		}
	}
	return
}

func RegisteUser(data *model.RegistData) error {
	role := 0
	//TODO 添加管理员检测
	_, err := db.Exec("insert into user(account,pwd,msg_tmp1) values(?,?,?)", data.Account, data.Pwd, role)
	return err
}

func DelUser(account string) (err error) {
	_, err = db.Exec("delete from user where account = ?", account)
	return
}
