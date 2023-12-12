package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func IsDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func Btoi(b []byte) int {
	n, _ := strconv.Atoi(string(b))
	return n
}

func Number(buf []byte, idx int) int {
	beg, end := idx, idx

	for IsDigit(buf[idx-1]) {
		idx -= 1
	}
	beg = idx
	for IsDigit(buf[idx+1]) {
		idx += 1
	}
	end = idx + 1

	return Btoi(buf[beg:end])
}

func Schema(filename string) (schema []byte, width int, height int) {
	buf, _ := os.ReadFile(filename)
	n := len(buf)

	width = bytes.IndexByte(buf, '\n')
	height = n / (width + 1)

	schema = make([]byte, (height+2)*(width+2))

	for i := 0; i < width+2; i += 1 {
		schema[i] = '.'
	}
	for i := 1; i <= height; i += 1 {
		schema[i*(width+2)] = '.'
		copy(schema[i*(width+2)+1:][:width], buf[(i-1)*(width+1):][:width])
		schema[i*(width+2)+width+1] = '.'
	}
	for i, end := 0, schema[(height+1)*(width+2):]; i < width+2; i += 1 {
		end[i] = '.'
	}

	return schema, width, height
}

func main() {
	filename := flag.String("i", "input", "input file")
	flag.Parse()
	schema, width, height := Schema(*filename)

	var numbers []int
	sum := 0
	for i := 1; i <= height; i += 1 {
		for j := 1; j <= width; j += 1 {
			if schema[i*(width+2)+j] != '*' {
				continue
			}

			left := i*(width+2) + j - 1
			right := i*(width+2) + j + 1
			top := (i-1)*(width+2) + j
			bottom := (i+1)*(width+2) + j

			if IsDigit(schema[left]) {
				numbers = append(numbers, Number(schema, left))
			}
			if IsDigit(schema[right]) {
				numbers = append(numbers, Number(schema, right))
			}
			if IsDigit(schema[top]) {
				numbers = append(numbers, Number(schema, top))
			} else {
				topLeft := top - 1
				topRight := top + 1

				if IsDigit(schema[topLeft]) {
					numbers = append(numbers, Number(schema, topLeft))
				}
				if IsDigit(schema[topRight]) {
					numbers = append(numbers, Number(schema, topRight))
				}
			}
			if IsDigit(schema[bottom]) {
				numbers = append(numbers, Number(schema, bottom))
			} else {
				bottomLeft := bottom - 1
				bottomRight := bottom + 1

				if IsDigit(schema[bottomLeft]) {
					numbers = append(numbers, Number(schema, bottomLeft))
				}
				if IsDigit(schema[bottomRight]) {
					numbers = append(numbers, Number(schema, bottomRight))
				}
			}
			if len(numbers) == 2 {
				sum += numbers[0] * numbers[1]
			}
			numbers = numbers[:0]
		}
	}

	fmt.Println(sum)
}
