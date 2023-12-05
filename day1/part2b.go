package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	engine := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")

	sum := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		str := scanner.Text()
		first, foundFirst := 0, false
		last := 0

		for {
			idxs := engine.FindStringIndex(str)
			if idxs == nil {
				break
			}
			digit := str[idxs[0]:idxs[1]]
			last = int(DigitValue[digit])
			if !foundFirst {
				first = last
				foundFirst = true
			}
			if len(digit) > 1 {
				str = str[idxs[1]-1:]
			} else {
				str = str[1:]
			}
		}
		sum += 10*first + last
	}
	fmt.Println(sum)
}

var DigitValue = map[string]int8{
	"1": 1, "one": 1,
	"2": 2, "two": 2,
	"3": 3, "three": 3,
	"4": 4, "four": 4,
	"5": 5, "five": 5,
	"6": 6, "six": 6,
	"7": 7, "seven": 7,
	"8": 8, "eight": 8,
	"9": 9, "nine": 9,
}
