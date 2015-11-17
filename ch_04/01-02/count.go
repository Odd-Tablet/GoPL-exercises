package main

import (
	"crypto/sha256"
	"fmt"
	"playground/GoPL-exercises/ch_04/01-02/popcount"
)

func main() {
	a := "a"
	b := "b"
	fmt.Println(a)
	h := sha256.New()
	h.Write([]byte(a))
	as := h.Sum(nil)
	h = sha256.New()
	h.Write([]byte(b))
	bs := h.Sum(nil)
	fmt.Println(b)
	fmt.Printf("%x\n", as)
	fmt.Printf("%x\n", bs)
	fmt.Printf("Difference in bits: %d\n", shaComp(as, bs))

}

func shaComp(a, b []byte) int {
	count := 0
	alen := len(a)
	for i := 0; i < alen; i++ {
		x := a[i] ^ b[i]
		count += popcount.PopCount(x)
	}
	return count
}
