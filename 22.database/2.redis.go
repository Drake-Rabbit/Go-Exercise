package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"

	"log"
)

// 我是单行注释

// var connErr error
var cRedis redis.Conn
var reErr error

func init() {

	cRedis, reErr = redis.Dial("tcp", "localhost:6379")
	if reErr != nil {
		log.Fatal(reErr.Error())
	}
}

// 存入一个数据
func doString() {
	cRedis.Do("set", "age", 66)
	cRedis.Do("expire", "age", 10) //设置age过期时间
	age, _ := redis.String(cRedis.Do("get", "age"))
	fmt.Println(age)
}

// 存入多数据
func doMString() {
	cRedis.Do("mset", "userName", "lampol", "pass", "123456")
	user, _ := redis.Strings(cRedis.Do("mget", "userName", "pass"))
	fmt.Println(user)
}

// 列表
func doList() {
	cRedis.Do("lpush", "num", 3, 6, 9)
	num, _ := redis.Int(cRedis.Do("rpop", "num"))
	fmt.Println(num)
}

// 哈希
func doHash() {
	cRedis.Do("hset", "userinfo", "userName", "lampol")
	//cRedis.Do("hset", "userinfo", "userName", "xxxx")
	str, _ := redis.String(cRedis.Do("hget", "userinfo", "userName"))
	fmt.Println(str)
}

func main() {
	//init()
	//doString()
	//doMString()
	//doList()
	doHash()
}

/*
多行注释
*/
