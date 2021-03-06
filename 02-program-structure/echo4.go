package tmpconv

import (
	"flag"
	"fmt"
	"strings"
)

// n and sep and pointers
var n = flag.Bool("n", false, "commit trailing newline")
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	result := strings.Join(flag.Args(), *sep)
	fmt.Print(result)

	if !*n {
		fmt.Println()
	}
}
