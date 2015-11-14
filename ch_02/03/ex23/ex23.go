// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package ex23

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
// Modified to a loop in accordance with exercise 2.3
func PopCount(x uint64) int {
	var result byte
	for x > 0 {
		result += pc[byte(x)]
		x = x >> 8
	}
	return int(result)
}

//!-
