package main

import "fmt"

//函数可以返回多个值
func text() (a,b,c int) {
	return 1, 2, 3
}

func main() {
	//多重赋值
	a,b,c:=10,20,30
	b,a = a,b//变量交换
	fmt.Println(a,b,c)

	//_匿名变量  丢弃数据不处理
	tmp,_ := 7,8
	fmt.Println(tmp)
	//当我们不想全部接受返回值时，这时使用推导式就行
	_,e,f := text()
	fmt.Println(e,f)
}
