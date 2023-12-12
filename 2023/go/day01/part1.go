package main

import (
	"bufio"
	"fmt"
	"os"
)

func isASCIIDigit(r rune) bool {
	if r < 0x80 && r >= '0' && r <= '9' {
		return true
	}
	return false
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	sum := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		first, firstFound := 0, false
		last := 0

		for _, c := range scanner.Text() {
			if isASCIIDigit(c) {
				last = int(c - '0')
				if !firstFound {
					first = last
					firstFound = true
				}
			}
		}

		sum += 10*first + last
	}

	fmt.Println(sum)
}
