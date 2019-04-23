package main

import "fmt"

func demo(s []int) []int {
	fmt.Println(s)
	s[1] = 123
	//如果函数中使用append进行数据添加  形参地址发生改变   就会指向实参地址
	s = append(s,6,7,87)
	fmt.Println(s)
	return s
}


func main() {
	slice := [] int{1,2,3,4,5}
	//地址传递   切片名是一个地址
	fmt.Println(slice)
	slice = demo(slice)
	fmt.Println(slice)
}
