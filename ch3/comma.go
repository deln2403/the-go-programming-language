// insert commas every three places in a string representation of an integer
// e.g. 12345 -> 12,345
package main

import (
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	for _, value := range os.Args[1:] {
		fmt.Println(comma(value))
	}
}