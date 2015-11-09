package main

import (
	"fmt"
	"os"
	"strings"
)

// Modify echo to print os.Args[0]

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
