package main

import "fmt"

func main1001() {
	var s [] int = [] int{1,2,3,6,5}
	s1:=make([] int,8)
	//使用copy  在内存中存储俩个独立的切片，一个发生修改不会改变另一个
	copy(s1,s)
	s1[4] = 9999
	fmt.Println(s1)
	fmt.Println(s)

	s1 = append(s1,9,8)
	fmt.Println(s1)
	fmt.Println(s)
	copy(s,s1)  //如果s1中的长度比s的长度大，直接截取长度为s1的部分覆盖s
	fmt.Println(s1)
	fmt.Println(s)

}
func main() {
	var s [] int = [] int{1,3,2,6,5}
	for i:=0;i<len(s);i++{
		for j:=0;j<i;j++{
			if s[i]>s[j]{
				s[i],s[j] = s[j],s[i]
			}
		}
	}
	fmt.Println(s)
}
