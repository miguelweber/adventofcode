package main

// TODO: optimize this slow piece of shit

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"slices"
)

type Conversion struct {
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

	seedRanges, advance := GetNumbers(buf, []uint32{})

	const nOfConversions = 7
	conversions := [nOfConversions][]Conversion{}

	numbers := []uint32{}
	for i := 0; i < nOfConversions; i++ {
		buf = buf[advance:]
		numbers, advance = GetNumbers(buf, numbers)

		for j := 0; j+3 <= len(numbers); j += 3 {
			conversions[i] = append(conversions[i], Conversion{
				src:   numbers[j+1],
				len:   numbers[j+2],
				delta: numbers[j+0] - numbers[j+1],
			})
		}
		slices.SortFunc(conversions[i], func(c1, c2 Conversion) int {
			if c1.src < c2.src {
				return -1
			}
			return 1
		})
	}
	buf, numbers = nil, nil // GC 'em

	result := ^uint32(0)
	for i := 0; i+2 <= len(seedRanges); i += 2 {
		var seeds []uint32
		for j := uint32(0); j < seedRanges[i+1]; j++ {
			seeds = append(seeds, seedRanges[i+0]+j)
		}
		for _, table := range conversions {
			for j, seed := range seeds {
				k, found := slices.BinarySearchFunc(table, seeds[j], func(r Conversion, n uint32) int {
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
					seeds[j] = seed + table[k].delta
				}
			}
		}
		smallestSeed := slices.Min(seeds)
		if smallestSeed < result {
			result = smallestSeed
		}
	}
	fmt.Println(result)
}
