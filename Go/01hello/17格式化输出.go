package main

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := "a"
	d := 3.2234
	//%T打印字符类型
	fmt.Printf("%T,%T,%T,%T\n",a,b,c,d)
	//%d 整型输出
	//%c 字符输出
	//%s 字符串输出
	//%f 浮点型输出
	fmt.Printf("%d,%s,%s,%.2f\n",a,b,c,d)

	//%v 自动格式匹配的输出
	fmt.Printf("%v,%v,%v,%v",a,b,c,d)


}
