package main

import "fmt"

type Attribute struct {
	Color string
	Size  float32
}

// 我是单行注释
func main() {
	////
	var i interface{}
	i = 4
	fmt.Printf("%T", i)
	var b int
	//b = i
	b = i.(int)
	fmt.Println(b)

	////
	goodsInfo := make(map[string]interface{})
	goodsInfo["name"] = "huawei"
	goodsInfo["num"] = 56
	goodsInfo["img"] = []string{"1", "2"}
	fmt.Println(goodsInfo["name"])
	fmt.Println(goodsInfo["num"])
	fmt.Println(goodsInfo["img"])
	//fmt.Println(goodsInfo["img"][0])  //访问不了的
	img, _ := goodsInfo["img"].([]string)
	fmt.Println(img[0])
	fmt.Println(goodsInfo["img"].([]string)[0]) //这里使用了接口的断言

	attr := Attribute{
		"red",
		5.6,
	}
	goodsInfo["attr"] = attr
	fmt.Println(goodsInfo["attr"])
	//fmt.Println(goodsInfo["attr"].Size) 访问不了的
	attr2, _ := goodsInfo["attr"].(Attribute) //这里使用了接口的断言
	fmt.Println(attr2.Color)
}

/*
多行注释
*/
