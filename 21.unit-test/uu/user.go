package uu

func Ssum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}
func Login(name, pass string) bool {
	if name == "map" && pass == "123" {
		return true
	}
	return false
}
