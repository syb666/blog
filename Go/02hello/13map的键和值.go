package main

import "fmt"

func main() {
	m:=make(map[string]string)
	m["sun"] = "sun"
	m["ya"] = "ya"
	m["bo"] = "bo"
	//delete(m["ya"])

	fmt.Println(m)
}
