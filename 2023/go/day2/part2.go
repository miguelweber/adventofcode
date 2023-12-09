package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Assumes input is always well-formed
func ParseGame(line string) int {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		switch r {
		case ':', ',', ' ':
			return true
		}
		return false
	})
	fields = fields[2:]

	maxRed, maxGreen, maxBlue := -1, -1, -1
	for lastSet := false; !lastSet; {
		for lastColor := false; !lastColor; {
			color := fields[1]
			if color[len(color)-1] == ';' {
				lastColor = true
				color = color[:len(color)-1]
			}
			n, _ := strconv.Atoi(fields[0])
			switch color {
			case "red":
				if n > maxRed {
					maxRed = n
				}
			case "blue":
				if n > maxBlue {
					maxBlue = n
				}
			default: // green
				if n > maxGreen {
					maxGreen = n
				}
			}

			fields = fields[2:]
			if len(fields) == 0 {
				lastColor = true
				lastSet = true
			}
		}
	}
	return maxRed * maxGreen * maxBlue
}

func main() {
	file, _ := os.Open("input")

	sum := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		text := scanner.Text()

		sum += ParseGame(text)
	}

	fmt.Println(sum)
}
