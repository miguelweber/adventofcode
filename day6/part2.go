package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

// Ignores spaces and stops when it encounters a newline
// Assumes input is >=0
// 1   23 4 -> 1234
func Btoi(bs []byte) int {
	result := 0

	for _, b := range bs {
		if b == '\n' {
			break
		}
		if b == ' ' {
			continue
		}
		result = result*10 + int(b) - '0'
	}

	return result
}

func main() {
	filename := flag.String("i", "input", "input file")
	flag.Parse()

	buf, _ := os.ReadFile(*filename)

	begTimeIdx := bytes.IndexByte(buf, ':') + 1
	begRecordIdx := bytes.IndexByte(buf[begTimeIdx:], ':') + begTimeIdx + 1

	time := Btoi(buf[begTimeIdx:])
	record := Btoi(buf[begRecordIdx:])

	ways := 0
	for i := 1; i < time; i++ {
		if i*(time-i) > record {
			ways++
		}
	}

	fmt.Println(ways)
}
