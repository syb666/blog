package main

import "fmt"

func main1() {
	var slice [] int
	//切片名放在栈区   数据放在堆区
	fmt.Printf("%p\n",slice)
	//使用append追加时地址可能会发生改变
	slice = append(slice,1,2,3,5)
	fmt.Printf("%p\n",slice)
	fmt.Printf("%p",&slice[1])

}
func main() {
	//容量
	s:=make([] int,5,10)
	s = append(s,1,2,3)
	fmt.Printf("%d-%d\n",len(s),cap(s))//长度-容量
	s = append(s,1,2,3)
	fmt.Printf("%d-%d",len(s),cap(s))//长度-容量
	//当数据超过原有容量时，一般容量以2倍的大小扩容   超过1024，每次扩容为原有的1/4
}
