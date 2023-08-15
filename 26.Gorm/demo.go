package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB
var dbErr error

func init() {
	dsn := "root:123456@tcp(192.168.1.10:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 创建的表名为tables不加s
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			//表前缀
			TablePrefix: "shop_",
		},
	})
	if dbErr != nil {
		log.Fatal(dbErr.Error())
	}
	fmt.Println(db)
}

type User struct {
	//gorm.Model
	Id         int    `gorm:"type:int;primary key"`
	Name       string `gorm:"type:varchar(20);not null ;index"`
	Age        int    `gorm:"type:tinyint;ungigned;not null"`
	Email      string `gorm:"type:varchar(20);index"`
	CreateTime int64  `gorm:"type:int;ungigned;column:create_at;autoCreateTime"`
	UpdateTime int64  `gorm:"type:int;ungigned;default:0;autoUpdateTime"`
}

// 创建的表名为tables不加s
//
//	func (u *User) TableName() string {
//		return "user"
//	}

// Hook 钩子函数

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	return
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
func (u *User) AfterSave(tx *gorm.DB) (err error) {
	return
}
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// 顺序
// BeforeSave
// BeforeCreate
// AfterSave
// AfterCreate
func main() {
	tx := db.Begin()
	user := User{Name: "lampo", Age: 34, Email: "werwe@qq.com"}
	//删除一个用户,创建一个用户，如果都没执行成功，则回滚
	delNum := tx.Delete(&User{}, 100).RowsAffected
	insNUM := tx.Create(&user).RowsAffected
	if delNum != 1 || insNUM != 1 {
		fmt.Println("roll back")
		tx.Rollback()
		return
	}
	tx.Commit()

	//type Result struct {
	//	Name string
	//	Age  int
	//}
	//var result Result
	//db.Raw("SELECT name,age FROM shop_user WHERE id= ?", 12).Scan(&result)
	//fmt.Println(result)
	//原生sql
	//res := db.Exec("INSERT INTO shop_user (name,age,email) VALUES (?,?,?)", "lampol_222", 22, "132465@qq.com")
	//res := db.Exec("UPDATE shop_user SET age=? WHERE id=?", 5, 12)
	//res := db.Exec("UPDATE shop_user SET age=? WHERE id=?", gorm.Expr("age+?", 10), 12)
	//fmt.Println(res.Error)
	//db.delete()
	//批量更新
	//res := db.Table("shop_user").Where("id=?", 6).Updates(&User{Name: "qwertyuiop", Age: 99})
	//fmt.Println(res.RowsAffected, res.Error)
	//update更新
	//res := db.Model(&User{}).Where("id=?", 5).Update("age", 66)
	//res := db.Table("shop_user").Where("id=?", 5).Update("age", 66)
	//fmt.Println(res.RowsAffected, res.Error)
	//join 联表查询
	type UserOrder struct {
		Name      string
		GoodsName string
	}
	//var userOrder UserOrder
	//db.Table("shop_user").Select("shop_user.name", "order_list.goods_name").
	//	Joins("join order_list on shop_user.id=order_list.user_id").Scan(&userOrder)
	//fmt.Println(userOrder)
	//数量count
	//var count int64
	//res := db.Table("shop_user").Where("age", 23).Count(&count)
	//fmt.Println(count, res.Error)
	//条件查询  =  <>  in  between and  lik….
	//不按主键   select , where  order, limit
	//var users []User
	//var user User

	//db.Where("age=?", 36).Find(&user)
	//db.Where("name=?", "lampo_9").Find(&user).Where("age=?", 36)
	//fmt.Println("user:", user, "\n", users)
	//db.Take(&user, map[string]interface{}{"name": "lampo_9", "age": 23})
	//fmt.Println(user)
	//db.Find(&users, "name=?", "lampo_3")
	//db.Take(&user, "name=?", "lampo_4")
	//fmt.Println(users, user)

	//只局部测试sql语句
	//stmt := db.Session(&gorm.Session{DryRun: true}).Take(&user, User{Name: "lampo_8", Age: 23}).Statement
	//fmt.Println(db.Dialector.Explain(stmt.SQL.String(), stmt.Vars))

	//
	//获取多个数据 Find 和 Scan
	//var users []User
	////db.Find(&users, []int{3, 4, 5})
	//db.Model(&users).Scan(&users)
	//fmt.Println(users)

	//根据主键检索数据
	//var user User
	//
	//db.Take(&user, 6)
	//fmt.Println(user)

	////Map查询
	//var data = map[string]interface{}{}
	////db.Model(&User{}).Last(&data)
	//
	////db.Table 只能使用table
	//db.Table("shop_user").Take(&data)
	//fmt.Println(data)

	//查询;GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象。
	//var userSelect User
	//// 获取第一条记录（主键升序）
	//db.First(&userSelect)
	//// 获取一条记录，没有指定排序字段
	//db.Take(&userSelect)
	//// 获取最后一条记录（主键降序）
	//db.Last(&userSelect)

	//指定表名查询数据
	//db.Table("shop_user").Last(&userSelect)
	//fmt.Println(userSelect)

	//map格式插入
	//var usermap map[string]interface{}
	//usermap = make(map[string]interface{})
	//usermap["name"] = "lampol_7"
	//usermap["email"] = "32434@qq.com"
	//usermap["age"] = 36
	//usermap["create_at"] = time.Now().Unix()
	//resmap := db.Model(&User{}).Create(usermap)
	//fmt.Println(resmap.RowsAffected, resmap.Error)
	//
	//map批量插入
	//create_at := time.Now().Unix()
	//usersMaps := []map[string]interface{}{
	//	{"Name": "lampo_8", "Age": 23, "Email": "63152qq.com", "create_at": create_at},
	//	{"Name": "lampo_9", "Age": 23, "Email": "63152qq.com", "create_at": create_at},
	//	{"Name": "lampo_10", "Age": 23, "Email": "63152qq.com", "create_at": create_at},
	//}
	//resMaps := db.Model(&User{}).Create(&usersMaps)
	//fmt.Println(resMaps.RowsAffected, resMaps.Error)
	//

	//批量插入
	//userslice1 := []User{
	//	{Name: "lampo_4", Age: 23, Email: "63152qq.com"},
	//	{Name: "lampo_5", Age: 23, Email: "63152qq.com"},
	//	{Name: "lampo_6", Age: 23, Email: "63152qq.com"},
	//}
	//res := db.Create(&userslice1)
	//fmt.Println(res.RowsAffected, res.Error)
	//	//自动迁移仅仅会创建表，添加缺少列和索引，并且不会改变现有列的类型或删除未使用的列
	//	//db.AutoMigrate(&User{})
	//	user := User{Name: "lampo_3", Age: 23, Email: "63152qq.com"}
	//	//res := db.Create(&user)
	//	//fmt.Println(res.RowsAffected, res.Error, user.Id)
	//	//选择性插入1.
	//	res2 := db.Select("name", "age").Create(&user)
	//	fmt.Println(res2.RowsAffected, res2.Error, user.Id)
	//	//选择性插入2.Omit,除了Email不插入
	//	//res3 := db.Omit("emali").Create(&user)
	//	//fmt.Println(res3.RowsAffected, res3.Error, user.Id)
}
