package main

import "fmt"

func main()  {
	var arr = [5] int{2,4,6,3,8}
	max1 := arr[0]
	min1 := arr[0]
	sum:=0
	for _,i:=range arr{
		if i>max1{
			max1 = i
		}
		if i<min1{
			min1 = i
		}
		sum+=i
	}
	length:=len(arr)
	fmt.Println(max1,min1,float64(sum)/float64(length))


	for j:=len(arr)-1;j>=0;j--{
		fmt.Println(arr[j])
	}



}
