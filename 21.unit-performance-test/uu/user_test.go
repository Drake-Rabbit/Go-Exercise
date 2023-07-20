package uu

import (
	"testing"
)

// 基准测试   性能比较
func benchmarkSsum(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Ssum(10000)
	}
}
func BenchmarkSsum1000(b *testing.B) {
	benchmarkSsum(b, 1000)
}

func BenchmarkSsum999999(b *testing.B) {
	benchmarkSsum(b, 999999)
}

// 单元测试
func TestSsum(t *testing.T) {
	//if testing.Short() {
	//	t.Skip("跳过这个测试函数")
	//}
	//
	//res := Ssum(5)
	////fmt.Println(res)
	//want := 10
	//if want != res {
	//	t.Error(want, res)
	//}
	t.Run("sum1", func(t *testing.T) {
		res := Ssum(5)
		//fmt.Println(res)
		want := 10
		if want != res {
			t.Error(want, res)
		}

	})
	t.Run("sum2", func(t *testing.T) {
		res := Ssum(5)
		//fmt.Println(res)
		want := 100
		if want != res {
			t.Error(want, res)
		}

	})

}

func TestLogin(t *testing.T) {
	res := Login("map", "123")
	if !res {
		t.Error(res)
	}
}
