package main

import "fmt"

func main() {
	//浮点型数据分为单精度 float32和双精度float64  精度不同
	var a float64 = 10.2389898945
	fmt.Println(a)

	var b float32 = 10.2389898945
	fmt.Println(b)
	//2.如果通过自动推导创建的， 默认是float64
	c := 3.14
	fmt.Println(c)
	fmt.Printf("%T",c)
}
