package main

import (
	"fmt"
	"os"
	"playground/GoPL-exercises/ch_02/01-02/tempconv"
	"strconv"
)

func convert(fl float64) {
	f := tempconv.Fahrenheit(fl)
	c := tempconv.Celsius(fl)
	k := tempconv.Kelvin(fl)
	fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c), c, tempconv.CToK(c), k, tempconv.KToC(k), k, tempconv.KToF(k), f, tempconv.FToK(f))
}

func main() {
	alen := len(os.Args)
	if alen > 1 {
		for _, arg := range os.Args[1:] {
			f, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("Usage: ./conv [number]")
				os.Exit(1)
			}
			convert(f)
		}
	} else {
		fmt.Printf("Enter a number to be converted: ")
		var arg float64
		_, err := fmt.Scanf("%f", &arg)
		if err != nil {
			fmt.Printf("Error scanning argument from stdin: ")
			panic(err)
		} else {
			convert(arg)
		}
	}
}
