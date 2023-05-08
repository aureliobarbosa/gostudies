package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	// var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	standardRegexpFilter := `^[a-z]+\[[0-9]+\]$`

	var validID = regexp.MustCompile(standardRegexpFilter)

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg, " ", validID.MatchString(arg))
	}
}
