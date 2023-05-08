package main

import (
	"fmt"
	"os"
)

/*
teste bla bla
*/

// Teste de comentário

func main() {
	var s, sep string

	fmt.Println(os.Args)

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i)
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
