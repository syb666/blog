package main

import "02hello/src/userinfo"

func main(){
	text()
	//包名.函数名  （函数首字母必须大写，否则找不到）
	userinfo.UserLogin()
}
