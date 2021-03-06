package tmpconv

import (
	"fmt"
	"os"
)

func main() {
	x, y := 1, 2
	x, y = y, x

	fmt.Printf("x = %d, y = %d\n", x, y) // x = 2, y = 1

	m1 := make(map[string]int)
	m1["x"] = 1
	m1["y"] = 2

	v, ok := m1["x"]
	if !ok {
		fmt.Fprintf(os.Stderr, "not exist in map")
	}

	fmt.Printf("v = %d\n", v)
}
