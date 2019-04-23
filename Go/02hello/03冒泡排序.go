package main

import "fmt"

func main() {
	var arr = [10] int {}
	var a int
	fmt.Scanf("%d",&a)
	for i:=0;i<a;i++{
		fmt.Scanf("%d",&arr[i])
	}

	for i:=0;i<a;i++{
		for j:=0;j<i;j++{
			if arr[i]>arr[j]{
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}

	for i:=0;i<a;i++{
		fmt.Println(arr[i])
	}
}
