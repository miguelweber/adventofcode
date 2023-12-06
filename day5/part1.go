package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"slices"
)

type Range struct {
	dst uint32
	src uint32
	len uint32
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

func Convert(n uint32, table []Range) uint32 {
	for i := range table {
		if n >= table[i].src && n < table[i].src+table[i].len {
			return n - table[i].src + table[i].dst
		}
	}
	return n
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
				dst: numbers[i+0],
				src: numbers[i+1],
				len: numbers[i+2],
			})
		}

		for i := range seeds {
			seeds[i] = Convert(seeds[i], table)
		}
	}

	fmt.Println(slices.Min(seeds))
}
