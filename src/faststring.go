package main

import (
	"os"
        "fmt"
        "strings"
        "strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	var b strings.Builder
        b.Grow(32)
	for i := 0; i < n; i++ {
		fmt.Fprintf(&b, "%d", i)
	}
        s := b.String()
	fmt.Println(s)
}
