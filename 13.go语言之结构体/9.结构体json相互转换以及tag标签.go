package main

import (
	"encoding/json"
	"fmt"
)

// 那我们定义结构体大写之后，
// 但是想让结构体中的字段 json 格式化为小写这时候就可以通过 tag 来指定：
type U9 struct {
	Name  string `json:"name"`
	Age   int
	Email string
}

func main() {

	user := U9{
		Name:  "lam",
		Age:   23,
		Email: "121@qq.com",
	}
	// 结构体 json相互转换
	fmt.Println(user)
	jsonByte, _ := json.Marshal(user) // 小写的属性不会显示为json
	data := string(jsonByte)
	fmt.Println(data)

	//json转换为结构体
	str := `{"Name":"lam","Age":23,"Email":"121@qq.com"}`
	var user2 U9
	err := json.Unmarshal([]byte(str), &user2) //返回nil则成功
	fmt.Printf("%#v,\n%T", err, err)

}

/*
多行注释
*/
