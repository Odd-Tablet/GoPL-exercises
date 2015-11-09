package main

import (
	"fmt"
	"os"
	"strconv"
)

// Modify echo to print index and value of arguments

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Index: %s Value: %s\n", strconv.Itoa(i), os.Args[i])
	}
}
