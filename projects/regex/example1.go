package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i], " ", validID.MatchString(os.Args[i]))
	}
}
