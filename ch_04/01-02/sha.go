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

func shaHash(str string) {
	reader := bufio.NewReader(os.Stdin)
	if str == "t" {
		fmt.Print("Enter value to be hashed in SHA384: ")
		a, _ := reader.ReadString('\n')
		a = strings.TrimRight(a, "\n")
		h := sha512.New384()
		h.Write([]byte(a))
		as := h.Sum(nil)
		fmt.Printf("%x\n", as)
	} else if str == "f" {
		fmt.Print("Enter value to be hashed in SHA512: ")
		a, _ := reader.ReadString('\n')
		a = strings.TrimRight(a, "\n")
		h := sha512.New()
		h.Write([]byte(a))
		as := h.Sum(nil)
		fmt.Printf("%x\n", as)
	} else {
		fmt.Print("Enter value to be hashed in SHA256: ")
		a, _ := reader.ReadString('\n')
		a = strings.TrimRight(a, "\n")
		h := sha256.New()
		h.Write([]byte(a))
		as := h.Sum(nil)
		fmt.Printf("%x\n", as)
	}
}
func main() {
	flag.Parse()
	if *tFlag == true {
		shaHash("t")
	} else if *fFlag == true {
		shaHash("f")
	} else {
		shaHash("c")
	}
}
