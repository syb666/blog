package main

import "fmt"

func main() {
	score:=160
	if score>100{
		fmt.Println("1")
		if score>150 && score<180{
			fmt.Println("2")
		}
	}else{
		fmt.Println("0")
	}
}
