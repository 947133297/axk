package util

import (
	"database/sql"
	"strconv"
	"strings"

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

	_, err = db.Exec(`
		create table if not exists project(
			id int(10) primary key not  null  auto_increment,
			uid int(10) ,
			name varchar(50) not null,
			CONSTRAINT pfk FOREIGN KEY (uid) REFERENCES user (id)
		)`)
	HandleErr(err, true)

	_, err = db.Exec(`
		create table if not exists section(
			id int(10) primary key not  null  auto_increment,
			pid int(10) ,
			uid int(10) ,
			name varchar(50) not null,
			graph varchar(50) not null,
			CONSTRAINT uufk FOREIGN KEY (uid) REFERENCES user (id),
			CONSTRAINT pfk2 FOREIGN KEY (pid) REFERENCES project (id)
		)`)
	HandleErr(err, true)
}

func UserLogin(data *model.LoginData) (user *model.User, err error) {
	return queryUser(true, -1, data.Account, data.Pwd)
}
func queryUser(isLogin bool, uid int, account, pwd string) (user *model.User, err error) {
	user = new(model.User)
	if isLogin {
		user.Account = account
		err = db.QueryRow("SELECT msg_tmp1,id FROM user WHERE account=? and pwd=?", account, pwd).Scan(&user.Role, &user.Id)
	} else {
		user.Id = uid
		err = db.QueryRow("select msg_tmp1,account from user where id=?", uid).Scan(&user.Role, &user.Account)
	}
	switch {
	case err == sql.ErrNoRows:
		user = nil
	case err != nil:
		Println(err.Error())
		user = nil
	}
	return
}

func GetUser(uid int) (user *model.User, err error) {
	return queryUser(false, uid, "", "")
}

func RegisteUser(data *model.RegistData) error {
	role := 0
	//TODO 添加管理员检测
	if strings.HasSuffix(data.Account, "999") {
		role = 1
	}
	r, err := db.Exec("insert into user(account,pwd,msg_tmp1) values(?,?,?)", data.Account, data.Pwd, role)
	i64, err := r.LastInsertId()
	if err != nil {
		return err
	}
	i32 := int(i64)
	return AddProject(i32, "默认项目")
}

func DelUser(account string) (err error) {
	_, err = db.Exec("delete from user where account = ?", account)
	return
}

func GetUserList() (list []*model.User) {
	rows, err := db.Query("SELECT id,msg_tmp1,account FROM user where msg_tmp1 = 0")
	if err != nil {
		Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		u := new(model.User)
		if err := rows.Scan(&u.Id, &u.Role, &u.Account); err != nil {
			Println(err.Error())
			return
		}
		list = append(list, u)
	}
	if err := rows.Err(); err != nil {
		Println(err.Error())
	}
	return
}

func AddProject(uid int, projectName string) error {
	_, err := db.Exec("insert into project(uid,name) values(?,?)", uid, projectName)
	return err
}

func GetAllProject(uid int) (list []*model.Project) {
	rows, err := db.Query("SELECT id,uid,name FROM project where uid =?", uid)
	if err != nil {
		Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := new(model.Project)
		if err := rows.Scan(&p.Id, &p.Uid, &p.Name); err != nil {
			Println(err.Error())
			return
		}
		list = append(list, p)
	}
	if err := rows.Err(); err != nil {
		Println(err.Error())
	}
	return
}

func GetProjectMsg(uid int, pid string) (p *model.Project) {
	p = new(model.Project)
	_pid, _ := strconv.Atoi(pid)
	p.Id, p.Uid = _pid, uid
	err := db.QueryRow("select name from project where id=? and uid=?", p.Id, p.Uid).Scan(&p.Name)
	switch {
	case err == sql.ErrNoRows:
		p = nil
	case err != nil:
		Println(err.Error())
		p = nil
	}
	return
}

func AddSection(data *model.Section) error {
	_, err := db.Exec("insert into section(uid,pid,name,graph) values(?,?,?,?)", data.Uid, data.Pid, data.Name, data.Graph)
	return err
}

func GetAllSection(pid string, uid int) (list []*model.Section) {
	rows, err := db.Query("SELECT id,uid,pid,name,graph FROM section where uid =? and pid = ?", uid, pid)
	if err != nil {
		Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		s := new(model.Section)
		if err := rows.Scan(&s.Id, &s.Uid, &s.Pid, &s.Name, &s.Graph); err != nil {
			Println(err.Error())
			return
		}
		list = append(list, s)
	}
	if err := rows.Err(); err != nil {
		Println(err.Error())
	}
	return
}
