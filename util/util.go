package util

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var logger *log.Logger

// InitLog 初始化日志
func InitLog(dirName, logName string) (err error) {
	dir, _ := os.Getwd()
	os.Mkdir(dir+"/"+dirName, os.ModePerm)
	logName = dirName + "/" + logName
	var file *os.File
	if _, err = os.Stat(logName); os.IsNotExist(err) {
		file, err = os.Create(logName)
	} else {
		file, err = os.OpenFile(logName, os.O_APPEND, 0644)
	}
	if err == nil {
		logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	}
	return
}

func Println(v ...interface{}) {
	logger.Println(v)
}

func Fatalln(v ...interface{}) {
	logger.Fatalln(v)
}

const _DB_NAME = "axk"

func init() {
	initDB()

}
func initDB() {
	// create db if not exists
	db, err := sql.Open("mysql", "root:qq5566@/mysql")
	HandleErr(err, true)
	db.Exec("create database " + _DB_NAME)
	db, err = sql.Open("mysql", "root:qq5566@/"+_DB_NAME)
	HandleErr(err, true)

	//create table # USER
	_, err = db.Exec(`
		create table if not exists user(
			id int(10),
			account varchar(30),
			pwd varchar(30) not null,
			msg_tmp1 varchar(50),
			msg_tmp2 varchar(50),
			PRIMARY KEY (id),
			unique key (account)
		)`)
	HandleErr(err, true)
}

func HandleErr(err error, ifExit bool) {
	if err != nil {
		if ifExit {
			Fatalln(err.Error())
		}
		Println(err.Error())
	}
}
