package main

import "fmt"

func main() {
	var x rune = 10
	var y byte = 255
	var z uintptr = 30

	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)

	var v1 = (0x1 << 2)
	fmt.Println(v1)
}
