package main

import "fmt"

func map1(m map[int]string)  {
	fmt.Println(m)
}

func main() {
	m := map[int]string{1:"张三",2:"李四"}
	map1(m)
	fmt.Println(m)

}
