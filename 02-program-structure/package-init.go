package main

import "fmt"

var a = b + c
var b = f() + 1
var c = 1

func f() int {
	return c + 1
}

func main() {
	fmt.Println(a) // 4
}

func init() {
	println("this is init one")
}

func init() {
	println("this is init two")
}
