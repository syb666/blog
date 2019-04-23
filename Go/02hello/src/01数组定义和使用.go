package main

import "fmt"

func main01()  {
	//数组定义
	var a [10] int
	fmt.Println(a)

	//注意
	//error
	//var arr int = 0
	//var a[n] int
	//数组赋值
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	for i:=0;i<10;i++{
		a[i] = i
	}
	fmt.Println(a)

	//匿名变量
	for _,v :=range a{
		fmt.Println(v)
	}

	//var a [10] float64 //默认为0
	//var a1 [10] string //默认为空字符
	//var a2 [10] bool //默认为false

}
func main() {

	//常见问题
	//数组长度  常量
	//数组下标  不能越界

	//数组初始化
	//1.全部初始化
	var a [5] int = [5]int{1,2,3,4,5}
	//如果俩个数组类型相同，个数相同
	aa:=a
	fmt.Println(aa)
	fmt.Println(a)
	//2.自动推导
	b:= [5] int{1,2,3,4,5}
	fmt.Println(b)
	//部分初始化
	c := [5] int {3,2,1}
	fmt.Println(c)

	//指定某一个位置初始化
	d:=[5] int {1:2,4:4}
	fmt.Println(d)

	//通过初始化确定长度
	f:=[...]int{1,2,3,4,5,65}
	fmt.Println(len(f))


}
