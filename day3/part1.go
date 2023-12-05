package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func Btoi(b []byte) int {
	n, _ := strconv.Atoi(string(b))
	return n
}

func IsDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
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

	sum := 0
	var number []byte
	for i := 1; i <= height; i += 1 {
		const (
			seeking = iota
			found
			skipping
		)
		state := seeking
		number = number[:0]

		for j := 1; j <= width; j += 1 {
			curLine := i * (width + 2)
			char := schema[curLine+j]

			switch state {
			case seeking:
				if IsDigit(char) {
					state = found
					j -= 1
					continue
				}
			case found:
				if !IsDigit(char) {
					state = seeking
					number = number[:0]
					continue
				}
				prevLine := (i - 1) * (width + 2)
				nextLine := (i + 1) * (width + 2)
				neighbours := [...]byte{
					schema[curLine+j-1],  // left
					schema[curLine+j+1],  // right
					schema[prevLine+j],   // up
					schema[nextLine+j],   // down
					schema[prevLine+j-1], // up-left
					schema[prevLine+j+1], // up-right
					schema[nextLine+j-1], // down-left
					schema[nextLine+j+1], // down-right
				}
				if bytes.ContainsAny(neighbours[:], "+@#%*$/-=&") {
					state = skipping
				}
				number = append(number, char)
			case skipping:
				if IsDigit(char) {
					number = append(number, char)
				} else {
					state = seeking
					sum += Btoi(number)
					number = number[:0]
				}
			}
		}
		if state == skipping {
			sum += Btoi(number)
		}
	}

	fmt.Println(sum)
}
