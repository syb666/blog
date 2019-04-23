package main

import "fmt"

func main() {
	count:=0
	for i:=0;i<10;i++{
		for j:=0;j<=i ;j++  {
			count++

		}
	}
	fmt.Println(count)
}
