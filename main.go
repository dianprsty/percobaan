package main

import "fmt"

func main() {
	var m map[string]int
	m:= make(map(string)int))

	fmt.Println("m :", m)

	if m == nil {
		fmt.Println("m is nil")
	}

	m["januari"] = 10000
	m["februari"] = 11000

	//delete(m, "januari")
	fmt.Println("m:", m)
}
