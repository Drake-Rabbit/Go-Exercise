package main

import (
	"os"
)

func main() {
	//4.
	//bc, _ = ioutil.ReadFile("user/user.txt")
	//fmt.Println(string(bc))

	//3.readAll
	file, _ := os.Open("user/user.txt")
	defer file.Close()
	//b := make([]byte, 1024)
	//b, _ = io.ReadAll(file)
	//fmt.Printf("%s", b)

	//2.缓冲只读,不用转切片，但是一行一行读取
	//reader := bufio.NewReader(file)
	//for {
	//	line, err := reader.ReadString('\n') //按‘ ’里的来分隔，delim是字符
	//
	//	if err == io.EOF {
	//		break
	//	} //不用转切片，但是一行一行读取
	//	fmt.Println(line)
	//
	//}

	//1.只读
	//tmp := make([]byte, 1024)
	//n, err := file.Read(tmp)
	//if err == io.EOF {
	//	fmt.Println("read over")
	//	return
	//}
	//fmt.Println(string(tmp[:n]))

}
