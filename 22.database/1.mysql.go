package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
var dbErr error

func init() {
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"

	//笔记本的数据库
	//db, dbErr = sql.Open("mysql", "root:123456@tcp(10.1.1.11:3306)/test") // 设置连接数据库的参数
	//家里的数据库
	db, dbErr = sql.Open("mysql", "root:123456@tcp(192.168.1.10:3306)/test?charset=utf8") // 设置连接数据库的参数
	if dbErr != nil {
		log.Fatal(dbErr.Error())
	}
	if dbErr = db.Ping(); nil != dbErr {
		log.Fatal("数据库链接失败: " + dbErr.Error())
	}
	db.SetMaxIdleConns(200)
	db.SetMaxIdleConns(10)
}

func inserData() {
	//res, err := db.Exec("INSERT INTO user (name,email)VALUES (?,?)", "map", "666@qq.com")
	//prepare
	pre, err := db.Prepare("INSERT INTO user (name,email)VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, _ := pre.Exec("map", "666@qq.com")
	id, _ := res.LastInsertId()  //id号
	row, _ := res.RowsAffected() //修改到的行的数量
	fmt.Println(id, row)
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:name`
	Email string `json:email`
}

func updateData() {
	ret, err := db.Exec("UPDATE user SET name=? WHERE id=?", "mapppp", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
	fmt.Println(ret.RowsAffected())
}

func DeleteData() {
	ret, err := db.Exec("DELETE FROM user WHERE id=?", 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret.RowsAffected())
}

func queryOne() {
	var user User
	row := db.QueryRow("SELECT * FROM user WHERE id=2")
	if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
func queryAll() {
	var user User
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	users := make([]User, 0)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}
	fmt.Println(users)
}

// 事务的模拟
func transData() {
	trans, _ := db.Begin()
	res, _ := trans.Exec("INSERT INTO user (name , email )VALUES (?,?)", "ZZZ", "4654@qq.com")
	addNum, _ := res.RowsAffected()
	ret, _ := trans.Exec("DELETE FROM user WHERE id=?", 1200)
	delNum, _ := ret.RowsAffected()
	if addNum == delNum {
		trans.Commit()
		fmt.Println("事务执行成功")
	} else {
		trans.Rollback()
		fmt.Println("事务执行失败")
	}
}
func main() {
	defer db.Close()
	//inserData()
	queryAll()
	//queryOne()
	//updateData()
	//DeleteData()
	transData()
	//
	queryAll()

}
