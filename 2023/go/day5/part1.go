package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"slices"
)

type Range struct {
	src   uint32
	len   uint32
	delta uint32
}

func Btou32(bs []byte) uint32 {
	result := uint32(0)

	for _, b := range bs {
		result = result*10 + uint32(b) - '0'
	}

	return result
}

func GetNumbers(buf []byte, backing []uint32) ([]uint32, int) {
	beg := bytes.IndexByte(buf, ':') + 2
	end := bytes.Index(buf, []byte("\n\n"))

	var advance int
	if end >= 0 {
		end += beg
		advance = end + 2
	} else {
		end = len(buf[beg:])
		advance = -1
	}

	begN := -1
	i := 0
	buf = buf[beg:]
	backing = backing[:0]
loop:
	for i = range buf {
		switch buf[i] {
		case '\n':
			if i+1 == end || buf[i+1] == '\n' {
				break loop
			}
			fallthrough
		case ' ':
			backing = append(backing, Btou32(buf[begN:i]))
			begN = -1
		default:
			if begN < 0 {
				begN = i
			}
		}
	}
	backing = append(backing, Btou32(buf[begN:i]))

	return backing, advance
}

func main() {
	filename := flag.String("i", "input", "input file")
	flag.Parse()
	buf, _ := os.ReadFile(*filename)

	seeds, advance := GetNumbers(buf, []uint32{})
	numbers := []uint32{}
	table := []Range{}

	for advance != -1 {
		buf = buf[advance:]
		numbers, advance = GetNumbers(buf, numbers)
		table = table[:0]

		for i := 0; i+3 <= len(numbers); i += 3 {
			table = append(table, Range{
				src:   numbers[i+1],
				len:   numbers[i+2],
				delta: numbers[i+0] - numbers[i+1],
			})
		}

		slices.SortFunc(table, func(r1, r2 Range) int {
			if r1.src < r2.src {
				return -1
			}
			return 1
		})

		for i, seed := range seeds {
			j, found := slices.BinarySearchFunc(table, seeds[i], func(r Range, n uint32) int {
				switch {
				case n < r.src:
					return 1
				case n >= r.src+r.len:
					return -1
				default:
					return 0
				}
			})
			if found {
				seeds[i] = seed + table[j].delta
			}
		}
	}

	fmt.Println(slices.Min(seeds))
}
