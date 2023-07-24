package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//	func init() {
//		var db *sql.DB
//		var dbErr error
//		//DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", root, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
//		db, dbErr = sql.Open("mysql", "root:123456@tcp(10.1.1.11:3306)/test")
//		if dbErr != nil {
//			log.Fatal(dbErr.Error())
//		}
//		if dbErr = db.Ping(); nil != dbErr {
//			log.Fatal("数据库链接失败: " + dbErr.Error())
//		}
//		db.SetMaxIdleConns(200)
//		db.SetMaxIdleConns(10)
//	}
func main() {
	//init()
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	var db *sql.DB
	var dbErr error
	db, dbErr = sql.Open("mysql", "root:@tcp(10.1.1.11:3306)/test") // 设置连接数据库的参数
	if dbErr != nil {
		log.Fatal(dbErr.Error())
	}
	if dbErr = db.Ping(); nil != dbErr {
		log.Fatal("数据库链接失败: " + dbErr.Error())
	}
	db.SetMaxIdleConns(200)
	db.SetMaxIdleConns(10)

}
