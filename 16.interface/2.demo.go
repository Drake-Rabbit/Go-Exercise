package main

import "fmt"

// 2.接口类型作为参数
type Mobiler interface {
	Call(int)
}
type Huawei struct {
	Name string
}
type Xiaomi struct {
	Name string
}

func (h Huawei) Call(mo int) {
	fmt.Println(h.Name, "get phone", mo)
}

func (h Xiaomi) Call(mo int) {
	fmt.Println(h.Name, "get phone", mo)
}

type Client struct {
}

func (c Client) run(m Mobiler, mo int) {
	//m.Call(mo)
	//断言判断 接口类型
	switch m.(type) {
	case Huawei:
		fmt.Println("调用华为的接口")

	case Xiaomi:
		fmt.Println("调用小米的接口")
	}
}
func main() {
	//var m1 Mobiler = Huawei{"P40"}
	//var m2 Mobiler = Xiaomi{"K60"}
	//
	//var c = Client{}
	//c.run(m1, 155755555555)
	//c.run(m2, 555554649849)

	//断言判断 接口类型
	//查询一个接口变量的底层变量是否还实现了其他接口
	var n1 Mobiler = Huawei{"P60"}
	var n2 Mobiler = Xiaomi{"K60"}
	var c1 = Client{}
	c1.run(n1, 231355)
	c1.run(n2, 54545454646)

}
