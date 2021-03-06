package main

import (
	"fmt"
	"math/rand"
)

func f() {}

var g = "g"

func main() {
	f := "f"
	fmt.Println(f) // "f"
	fmt.Println(g) // "g"

	toUpper1() // "HELLO"
	toUpper2() // "HELLO"

	printValue()
	printValue2()
}

func toUpper1() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println()
}

func toUpper2() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
	fmt.Println()
}

func printValue() {
	if x := rand.Intn(50); x > 10 {
		fmt.Println(x)
	} else if y := rand.Intn(50); y > x {
		fmt.Println(y)
	}
	// fmt.Printf("x = %d, y = %d\n", x, y)	// compile error: x and y are not visible here
}

func printValue2() {
	switch x := rand.Intn(10); x {
	case 5:
		fmt.Printf("x = %d\n", x*2)
	default:
		fmt.Printf("x = %d\n", x*3)
	}
}
