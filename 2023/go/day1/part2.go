package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Match(str string) (found bool, value, advance int) {
	for i := range str {
		view := str[i:]
		for j := range digits {
			if strings.HasPrefix(view, digits[j]) {
				var length int
				if len(digits[j]) > 1 {
					length = i + len(digits[j]) - 1
				} else {
					length = 1
				}

				return true, int(digitsValue[j]), length
			}
		}
	}
	return false, 0, 0
}

var digits = [...]string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9",
	"one", "two", "six",
	"four", "five", "nine",
	"three", "seven", "eight",
}
var digitsValue = [...]int8{
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 6,
	4, 5, 9,
	3, 7, 8,
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	sum := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		str := scanner.Text()
		first, foundFirst := 0, false
		last := 0

		for {
			found, val, advance := Match(str)
			if !found {
				break
			}
			last = val
			if !foundFirst {
				first = val
				foundFirst = true
			}
			str = str[advance:]
		}
		sum += 10*first + last
	}

	fmt.Println(sum)
}
