package main

import "fmt"

func main() {
	fmt.Println("Hiragana represented as: ")
	runa := 0x3040

	fmt.Printf("%32s %8s %8s %8s %4s\n", "Binary", "Octal", "Integer", "Hexadecimal", "Runa")
	for runa < 0x30a0 {
		fmt.Printf("%32b %8o %8d %8x %4c\n", runa, runa, runa, runa, runa)
		runa++
	}

	fmt.Println("")

}
