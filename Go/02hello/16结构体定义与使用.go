package main

import (
	"fmt"
	//"go/ast"
)

type student struct {
	id int
	name string
	score int
}

func Sort(arr []student) {
	for i:=0;i<len(arr);i++{
		for j:=0;j<i;j++{
			if arr[i].score>arr[j].score{
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
}

func main() {
	var arr [] student = []student{
		{id:1,name:"sunyabo1",score:90},
		{id:2,name:"sunyabo2",score:80},
		{id:3,name:"sunyabo3",score:70}}
	fmt.Println(arr)
	arr[0].id = 99
	for _,i := range arr{
		fmt.Println(i)
	}

	//结构体作为   值传递
	Sort(arr)
	fmt.Println(arr)

	arr = append(arr,student{id:6,name:"999",score:11})
	//切片作为函数参数   地址传递（引用）
	Sort(arr)
	fmt.Println(arr)
}
