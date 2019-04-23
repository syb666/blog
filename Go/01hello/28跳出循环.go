package main

import (
	"fmt"
)

func main() {
	//在循环里面有两个关键操作break和continue，break操作是跳出当前循环，continue是跳过本次循环。
	for i := 0; i < 5; i++ {
		//time.Sleep(time.Second)
		if 2 == i {
			//break//break操作是跳出当前循环
			continue//continue是跳过本次循环
		}
		fmt.Println(i)
	}

	//用goto跳转到必须在当前函数内定义的标签：
	for i := 0; i < 5; i++{
		for{
			fmt.Println(i)
			goto LABEL//跳转到标签LABEL，从标签处，执行代码
		}
	}
	fmt.Println("thisistest")
	LABEL:
		fmt.Println("itisover")


}
