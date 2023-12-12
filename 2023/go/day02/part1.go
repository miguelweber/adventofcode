package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

// Assumes input is always well-formed
func ValidGame(line string) (bool, int) {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		switch r {
		case ':', ',', ' ':
			return true
		}
		return false
	})
	gameNum, _ := strconv.Atoi(fields[1])
	fields = fields[2:]

	for lastSet := false; !lastSet; {
		var red, green, blue int

		for lastColor := false; !lastColor; {
			color := fields[1]
			if color[len(color)-1] == ';' {
				lastColor = true
				color = color[:len(color)-1]
			}
			n, _ := strconv.Atoi(fields[0])
			switch color {
			case "red":
				red = n
			case "blue":
				blue = n
			default: // blue
				green = n
			}
			fields = fields[2:]
			if len(fields) == 0 {
				lastColor = true
				lastSet = true
			}
		}
		if red > MaxRed || green > MaxGreen || blue > MaxBlue {
			return false, 0
		}
	}
	return true, gameNum
}

func main() {
	file, _ := os.Open("input")

	sum := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		text := scanner.Text()

		if ok, gameNum := ValidGame(text); ok {
			sum += gameNum
		}
	}

	fmt.Println(sum)
}
