package tmpconv

import "fmt"

func main() {
	p1 := new(int)
	fmt.Println(*p1) // 0

	*p1 = 1
	fmt.Println(*p1) // 1

	p2 := newIntByNew()
	p3 := newIntByVar()
	fmt.Println(p2 == p3) // false

	fmt.Println(delta(5, 10)) // 5
}

func newIntByNew() *int {
	return new(int)
}

func newIntByVar() *int {
	var dummy int
	return &dummy
}

func delta(old int, new int) int {
	return new - old
}
