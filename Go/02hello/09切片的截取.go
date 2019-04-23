package main

import "fmt"

func main() {
	s:= [] int{1,2,3,4,5,6,7}
	slice := s[1:3:5]
	fmt.Printf("%d  %d\n",len(slice),cap(slice))//2  4(5-1)


	//截取后切片的值还是原有的数组，改变切片值，会影响原有的数组
	s1:= [] int{1,2,3,4,5,6,7}
	slice1 := s1[1:3]
	slice1[1] = 9999
	fmt.Println(s1,slice1)//[1 2 9999 4 5 6 7]  [2 9999]
}
