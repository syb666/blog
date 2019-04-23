package main

import "fmt"

func main() {
	score:=200
	if score>100{
		fmt.Println("1")
	}else{
		fmt.Println("0")
	}

	if score>100{
		fmt.Println("1")
	}else if score>50{
		fmt.Println("0")
	}else if score>0 {
		fmt.Println("--0")
	}
}
