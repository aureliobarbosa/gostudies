package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0.0
	BoilingC      Celsius = 100.0
)

func (c Celsius) String() string    { return fmt.Sprintf("%4.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%4.2f°F", f) }

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "tempconv: %v\n", err)
		}

		f := Fahrenheit(t)
		c := Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n", f, FtoC(f), c, CtoF(c))
	}
	fmt.Printf("\n")
}
