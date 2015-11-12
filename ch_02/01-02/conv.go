package main

import (
	"./tempconv"
	"fmt"
	"os"
	"strconv"
)

func convert(fl float64) {

}

func main() {
	alen := len(os.Args)
	if alen >= 1 {
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
