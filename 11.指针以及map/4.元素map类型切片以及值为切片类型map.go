package main

import "fmt"

// 我是单行注释
func main() {
	//元素类型map的切片

	users := make([]map[string]string, 4)
	fmt.Println("元素切片map的本來面目:", users)
	users[0] = make(map[string]string, 3)
	users[1] = make(map[string]string, 3)
	//
	users[0]["name"] = "zyh"
	users[0]["age"] = "20"
	users[0]["email"] = "20@qq.com"
	//
	users[1]["name"] = "yph"
	users[1]["age"] = "20"
	users[1]["email"] = "2123@qq.com"

	fmt.Println(users)

	//值为切片map的切片
	//山东:{"青岛","济南","烟台"}  江苏:{"南京","苏州","常州")
	citys := make(map[string][]string, 1)
	//先定义一个切片
	sCity := make([]string, 3)
	sCity = append(sCity, "青岛", "济南", "烟台")
	citys["山东"] = sCity

	jCity := make([]string, 3)
	jCity = append(jCity, "南京", "苏州", "常州")
	citys["江苏"] = jCity

	fmt.Println()
	fmt.Println("值为切片map的本來面目:", citys)
	fmt.Println(citys)

}

/*
多行注释
*/
