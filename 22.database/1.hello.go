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
func main() {
	defer db.Close()
	inserData()

}
