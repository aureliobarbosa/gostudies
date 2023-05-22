package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Exemplo com inteiro:")
	fmt.Printf("%%d inteiro %d em formato decimal, math.MaxInt64 = %d\n", math.MaxInt64, math.MaxInt64)
	fmt.Printf("%%x inteiro %d em formato hexadecimal, math.MaxInt64 = %x\n", math.MaxInt64, math.MaxInt64)
	fmt.Printf("%%o inteiro %d em formato decimal, math.MaxInt64 = %o\n", math.MaxInt64, math.MaxInt64)
	fmt.Printf("%%b inteiro %d em formato decimal, math.MaxInt64 = %b\n", math.MaxInt64, math.MaxInt64)

	fmt.Println("\nExemplo com ponto flutuante:")
	fmt.Printf("%%f  ponto flutuante decimal, math.Pi = %f\n", math.Pi)
	fmt.Printf("%%g  ponto flutuante decimal, math.Pi = %g\n", math.Pi)
	fmt.Printf("%%e  ponto flutuante decimal, math.Pi = %e\n", math.Pi)

	fmt.Println("\nOutros exemplos:")
	fmt.Printf("%%c runa em formato hexadecimal, ぁ = %x\n", 'ぁ')
	fmt.Printf("%%c runa em unicode utf-8, %c\n", 'ぁ')

	runa := 0x3041

	fmt.Printf("%%c número hexadecimal 0x3041 como runa unicode utf-8, %c\n", runa)

}
