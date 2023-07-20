package uu

import (
	"testing"
)

func TestSsum(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过这个测试函数")
	}

	res := Ssum(5)
	//fmt.Println(res)
	want := 10
	if want != res {
		t.Error(want, res)
	}
}

func TestLogin(t *testing.T) {
	res := Login("map", "123")
	if !res {
		t.Error(res)
	}
}
