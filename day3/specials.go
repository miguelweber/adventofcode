package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	specials := make(map[byte]struct{})
	file, _ := os.Open("input")

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		text := scanner.Bytes()
		fmt.Println(len(text))
		for _, c := range text {
			if c != '.' && c != '\n' && !unicode.IsDigit(rune(c)) {
				specials[c] = struct{}{}
			}
		}
	}

	for k, _ := range specials {
		fmt.Printf("%c", rune(k))
	}
	fmt.Println()
}
