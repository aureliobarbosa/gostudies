package main

import (
	"fmt"
	"os"
)

func main() {

	text := "abcdefgABCDEFG"
	var data []byte
	data = []byte(text)

	fmt.Println(data)

	os.WriteFile("teste", data, 0666)
}
