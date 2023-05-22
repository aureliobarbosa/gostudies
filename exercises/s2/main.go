package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("You must provide the name of a file to be read!")
		os.Exit(1)
	}

	for i, filename := range os.Args {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			continue
		}
		fmt.Fprintf(os.Stdout, "%s:\n", filename)
		io.Copy(os.Stdout, file)
		fmt.Fprintf(os.Stdout, "\n")
		file.Close()
	}

}
