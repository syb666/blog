package main

import "fmt"

func main() {
	var str1 string
	str1  ="abc"
	fmt.Println(str1)

	//自动推导
	str2 :="孙亚博"
	fmt.Println(str2)

	//len是计算字符串的个数   不包含\0
	fmt.Println(len(str1))
	fmt.Println(len(str2))    //9   在Go中一个汉字占3个字符

	//字符串的拼接
	str3 := "hello"
	str4 := "孙亚"
	str5 :=str3+str4
	fmt.Println(str5)
}
