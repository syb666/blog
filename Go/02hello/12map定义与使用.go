package main

import "fmt"

func main() {
	var m map[int]string
	fmt.Println(m)

	m1 := make(map[int]string ,6)
	//键是唯一的，值不是   无序的
	m1[0] = "张三"
	m1[1] = "李四"
	m1[2] = "王二"
	m1[4] = "麻子"
	fmt.Println(m1)
	//初始化
	m3 := map[int]string{1:"张",2:"李",3:"王",4:"麻"}
	fmt.Printf("%p\n",m3)
	fmt.Println(m3)
	//map是自动扩容的

}
