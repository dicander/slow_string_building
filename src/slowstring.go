package main

import (
	"os"
        "fmt"
        "strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	answer := ""
	for i := 0; i < n; i++ {
		temp := strconv.Itoa(i)
		answer += temp
	}
	fmt.Println(answer)
}
