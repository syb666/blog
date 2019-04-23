package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机数种子一定的话，结果就是一样的
	rand.Seed(1)
	fmt.Println(rand.Int())//生成的数比较大

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))//生成[0-9]


	//双色球生成
	var red[6] int
	for i:=0;i<len(red);i++{
		v:= rand.Intn(33)+1
		for j:=0;j<=i;j++{
			if v==red[j] {
				v = rand.Intn(33)+1
				j=-1
			}
		}
		red[i] = v
	}
	fmt.Println(red)


}
