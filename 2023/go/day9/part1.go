package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("i", "input", "input file")
	flag.Parse()

	file, _ := os.Open(*filename)
	defer file.Close()

	numbers := []int32{}
	result := int32(0)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		numbers = GetNumbers(line, numbers)

		result += Extrapolate(numbers)
	}
	fmt.Println(result)
}

func Extrapolate(ns []int32) int32 {
	diffs := make([]int32, len(ns)-1)
	allZeros := true
	for i := 0; i+1 < len(ns); i++ {
		diffs[i] = ns[i+1] - ns[i+0]
		if diffs[i] != 0 {
			allZeros = false
		}
	}
	if allZeros {
		return ns[0]
	}

	return Extrapolate(diffs) + ns[len(ns)-1]
}

func GetNumbers(line string, buf []int32) []int32 {
	buf = buf[:0]

	result := int32(0)
	negative := int32(1)
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '-':
			negative = -1
		case ' ':
			buf = append(buf, result*negative)
			result = 0
			negative = 1
		default:
			result = result*10 + int32(line[i]) - '0'
		}
	}

	return append(buf, result*negative)
}
