package main

import "fmt"

func swap(a,b int) (int,int) {

	a,b = b,a
	return a,b
}

func sort(arr [5] int) ([5] int) {
	for i:=0;i<5;i++{
		for j:=0;j<i;j++{
			if arr[i]>arr[j]{
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
	//for i:=0;i<a;i++{
	//	fmt.Println(arr[i])
	//}
	return arr
}

func main() {

	a,b:=swap(2,3)
	fmt.Println(a,b)

	var arr = [5] int {4,3,2,8,1}
	//var a int
	//fmt.Scanf("%d",&a)
	//如果想通过函数计算结果改变实参的值，需要使用数组作为返回值
	arr1:=sort(arr)
	fmt.Println(arr1)


}
