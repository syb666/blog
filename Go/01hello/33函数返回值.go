package main

import "fmt"

func text001() (int,string)  {
	var a int = 10
	var b string = "1"
	return a,b
}



func main() {
	//var a int
	var b string
	//匿名变量就是丢弃不需要的数据
	_,b = text001()
	fmt.Println(b)
}
