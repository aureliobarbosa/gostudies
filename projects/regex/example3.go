package main

import (
	"flag"
	"fmt"
	"regexp"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	// var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	var regexpFilter string

	flag.StringVar(&regexpFilter, "regex", `^[a-z]+\[[0-9]+\]$`, "Defines a regular expression.")
	flag.Parse()

	fmt.Println("Using filter: ", regexpFilter)

	var validID = regexp.MustCompile(regexpFilter)

	for i, arg := range flag.Args() {
		fmt.Println(i, arg, " ", validID.MatchString(arg))
	}
}
