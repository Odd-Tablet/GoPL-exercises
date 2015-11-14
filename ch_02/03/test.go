package main

import (
	"fmt"
	"playground/GoPL-exercises/ch_02/03/ex23"
	"playground/GoPL-exercises/ch_02/03/popcount"
)

func main() {
	var x uint64
	for i := 0; i < 256; i++ {
		x = uint64(i)
		// Need to Benchmark results of looping and single execution
		fmt.Printf("Integer value: %d Byte Value: %d %d\n", x, popcount.PopCount(x), ex23.PopCount(x))
	}
}
