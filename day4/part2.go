package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

type scratchcard struct {
	matches, count int
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("Hi")

	file, _ := os.Open("bigboy.txt")
	scratchcards := []scratchcard{}

	for i, scanner := 0, bufio.NewScanner(file); scanner.Scan(); i++ {
		matches := CountMatches(scanner.Text())

		scratchcards = append(scratchcards, scratchcard{
			count:   1,
			matches: matches,
		})
	}
	sum := 0
	for i := range scratchcards {
		sum += scratchcards[i].count
		for j := 1; j <= scratchcards[i].matches; j++ {
			if i+j >= len(scratchcards) {
				break
			}
			scratchcards[i+j].count += scratchcards[i].count
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
