package main

import "fmt"

func main() {
	var a float64
	//阻塞进程 等待用户输入  格式化输入
	fmt.Scanf("%f",&a)

	//简单输入
	fmt.Scan(&a)
	fmt.Println(a)
}
