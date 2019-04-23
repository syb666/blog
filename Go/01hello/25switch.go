package main

import "fmt"

func main() {
	var score int = 90
	switch score {
	case 90:
		fmt.Println("A")
		//fallthrough
	case 80:
		fmt.Println("B")
		//fallthrough
	case 50,60,70:
		fmt.Println("C")
		//fallthrough
	default:
		fmt.Println("D")
	}

	switch s1:=90 ; s1{//初始化语句;条件
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	//2
	var s2 int=90
	switch{//这里没有写条件
	case s2 >= 90://这里写判断语句
		fmt.Println("A")
	case s2 >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	//3
	switch s3 := 90;{//只有初始化语句，没有条件
	case s3 >= 90://这里写判断语句
		fmt.Println("A")
	case s3 >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}
	/*
	优点
	1.if适合进行区间的判断   嵌套使用
	2.switch 固定值 执行效率高  可以将多个满足相同条件的值放在一起
	缺点
	1.if执行效率低
	2.switch 不建议使用嵌套
	*/
}
