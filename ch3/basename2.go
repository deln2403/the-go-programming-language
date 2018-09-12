package main

import (
	"fmt"
	"strings"
	"os"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	for _, value := range os.Args[1:] {
		fmt.Println(basename(value))
	}
}
