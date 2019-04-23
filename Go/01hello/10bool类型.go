package main
import "fmt"
func main() {
	//声明变量bool  没有初始化  就使用0
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	//2. bool不支持其他类型的赋值  不支持强制类型转换
	//a = 1
	//a = bool(1)
	//fmt.Println(a)
	//3. 自动推导
	var b = true
	fmt.Println(b)//代表代码冗余
	c:=true
	fmt.Println(c)

	v2 := (1==2)
	fmt.Println(v2)

}
