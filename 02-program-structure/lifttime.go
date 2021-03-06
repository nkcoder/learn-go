package tmpconv

import "fmt"

var global *int

func main() {
	global = f()

	g()

	fmt.Println(*global)
}

func f() *int {
	var x int
	x = 1
	return &x
}

func g() {
	y := new(int)
	*y = 1
}
