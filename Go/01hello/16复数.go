package main

import "fmt"

func main() {
	var a complex64
	a = 2.1 + 3.12i
	fmt.Println(a)

	//自动推导  默认complex128
	v2 := 3.2 + 12i
	fmt.Println(v2)

	fmt.Println(real(v2))//取实部
	fmt.Println(imag(v2))//取虚部


	v3 := complex(3.2, 12)
	fmt.Println(v3)
}
