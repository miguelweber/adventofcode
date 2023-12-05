package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("Hi")

	file, _ := os.Open("bigboy.txt")
	sum := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		matches := CountMatches(scanner.Text())

		if matches > 0 {
			sum += 1 << (matches - 1)
		}
	}
	log.Printf("sum: %d", sum)
}

func CountMatches(line string) int {
	words := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
	matches := 0
	winning := make(map[string]struct{})
	i := 0
	for i = range words {
		if words[i] == "|" {
			i += 1
			break
		}
		winning[words[i]] = struct{}{}
	}
	for _, num := range words[i:] {
		if _, ok := winning[num]; ok {
			matches++
		}
	}
	return matches
}
