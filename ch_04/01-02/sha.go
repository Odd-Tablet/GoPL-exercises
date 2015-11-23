package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

var tFlag = flag.Bool("t", false, "Sha384 Hash flag")
var fFlag = flag.Bool("f", false, "Sha512 Hash flag")

func shaHash(fl string, a string) {
	if fl == "t" {
		fmt.Print("SHA384 Hash Value:\n")
		fmt.Printf("%x\n", sha512.Sum512([]byte(a)))
	} else if fl == "f" {
		fmt.Print("SHA512 Hash Value:\n")
		fmt.Printf("%x\n", sha512.Sum512([]byte(a)))
	} else {
		fmt.Print("SHA256 Hash Value:\n")
		fmt.Printf("%x\n", sha256.Sum256([]byte(a)))
	}
}
func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value to be hashed: ")
	a, _ := reader.ReadString('\n')
	a = strings.TrimRight(a, "\n")
	if *tFlag == true {
		shaHash("t", a)
	} else if *fFlag == true {
		shaHash("f", a)
	} else {
		shaHash("c", a)
	}
}
