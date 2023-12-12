package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

// Assumes input is >=0
func Btoi(bs []byte) int {
	result := 0

	for _, b := range bs {
		result = result*10 + int(b) - '0'
	}

	return result
}

func IsNumeric(b byte) bool { return b >= '0' && b <= '9' }

func main() {
	filename := flag.String("i", "input", "input file")
	flag.Parse()

	buf, _ := os.ReadFile(*filename)

	var times, records []int

	doTimes := true
	for _, word := range bytes.Fields(buf)[1:] {
		if !IsNumeric(word[0]) {
			doTimes = false
			continue
		}
		if doTimes {
			times = append(times, Btoi(word))
		} else {
			records = append(records, Btoi(word))
		}
	}
	buf = nil

	result := 1

	for i := 0; i < len(times); i++ {
		ways := 0
		for j := 1; j < times[i]; j++ {
			if j*(times[i]-j) > records[i] {
				ways++
			}
		}
		result *= ways
	}

	fmt.Println(result)
}
