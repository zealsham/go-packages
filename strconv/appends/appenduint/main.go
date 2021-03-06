package main

import (
	"fmt"
	"strconv"
)

func main() {
	b10 := []byte("uint (base 10):")
	// AppendUint appends the string form of the unsigned integer i,
	// as generated by FormatUint, to dst and returns the extended buffer.
	b10 = strconv.AppendUint(b10, 15, 10)
	fmt.Println(string(b10)) // 15

	b16 := []byte("uint (base 16):")
	b16 = strconv.AppendUint(b16, 15, 16)
	fmt.Println(string(b16)) // f
}
