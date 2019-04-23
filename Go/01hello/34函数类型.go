package main

import "fmt"

func demo1()  {

}

type FUNC func()

func main() {
	//低地址
	var s = demo1
	var k FUNC
	k = demo1
	fmt.Println(demo1)
	fmt.Println(s)
	fmt.Println(k)
}
