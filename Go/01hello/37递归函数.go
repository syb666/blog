package main

import "fmt"

func demo(a int)  {
	if a<0{
		return
	}
	demo(a-1)
	fmt.Println(a)
}

func demo2(a int) int {
	if a<0{
		return 0
	}
	return a+demo2(a-1)
}

func main() {
	demo(4)
	fmt.Println(demo2(100))
}
