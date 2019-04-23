package main

import "fmt"

func text4() func() int {
	return func() int {
		x:=10
		return x
	}
}

func main() {
	var num int
	num = 9
	//匿名函数   没有名字的函数
	func(){
		num++
		fmt.Println(num)
	}()  //花括号  {}()函数调用

	f:= func() {
		num++
		fmt.Println(num)
	}

	type FuncType func()
	var f1 FuncType
	f1 = f
	f1()
	fmt.Println(num)

	//1. 参数传递
	func(a,b int){
		var sum int
		sum = a+b
		fmt.Println(sum)
	}(3,6) //调用时传值

	//匿名函数返回值
	x,y:= func(i,j,z int)(max,min int) {
		if i>j{
			max = i
			min = j
		}else{
			max = j
			min = i
		}
		return
	}(3,4,6)
	fmt.Println(x,y)


	//闭包
	f11:=text4()
	fmt.Println(f11())

}
