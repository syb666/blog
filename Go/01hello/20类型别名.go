package main

import "fmt"

func main() {
	//给int64起一个别名 bigint
	type bigint int64
	var a bigint
	fmt.Println(a)
	fmt.Printf("%T",a)

	type (
		long int64
		char byte
	)
	var b long = 11
	var ch char = 'a'
	fmt.Println(b,ch)
	fmt.Printf("%T-%T",b,ch)

}
