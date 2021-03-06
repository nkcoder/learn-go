package tmpconv

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Printf("*p = %d\n", *p) // *p = 1

	*p = 2
	fmt.Printf("*p = %d, x = %d\n", *p, x) // *p = 2, x = 2

	var p0 *int
	fmt.Println(p0 == nil) // true

	var i, j int
	fmt.Println(&i == &i, &i == &j, &i == nil) // true false false

	p1 := f()
	p2 := f()
	fmt.Printf("p1 == p2: %t, p1 = %d, p2 = %d\n", p1 == p2, *p1, *p2) // p1 == p2: false, p1 = 3, p2 = 3

	p3 := 1
	incr(&p3)
	fmt.Printf("p3 = %d\n", p3) // p3 = 2
}

func f() *int {
	v := 3
	return &v
}

func incr(p *int) {
	*p++
}
