package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

type Point struct{ x, y int }

type Galaxy struct {
	orig  Point
	delta Point
}

type Grid struct {
	points []byte
	Width  int
	Height int
}

func FindGalaxy(x, y int, gs []Galaxy) *Galaxy {
	point := Point{x, y}
	for i := range gs {
		if gs[i].orig == point {
			return &gs[i]
		}
	}
	return nil
}

func NewGrid(filename string) Grid {
	buf, err := os.ReadFile(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	width := bytes.IndexByte(buf, '\n')
	return Grid{
		buf,
		width,
		len(buf) / (width + 1),
	}
}

func (g *Grid) At(x, y int) byte { return g.points[y*(g.Width+1)+x] }

func main() {
	filename := flag.String("i", "input", "input file")
	flag.Parse()

	grid := NewGrid(*filename)
	galaxies := []Galaxy{}
	for y, expand_y := 0, 0; y < grid.Height; y++ {
		emptyLine := true
		for x := 0; x < grid.Width; x++ {
			if grid.At(x, y) == '#' {
				galaxies = append(galaxies, Galaxy{
					Point{x, y}, Point{0, expand_y},
				})
				emptyLine = false
			}
		}
		if emptyLine {
			expand_y++
		}
	}
	for x, expand_x := 0, 0; x < grid.Width; x++ {
		emptyRow := true
		for y := 0; y < grid.Height; y++ {
			if grid.At(x, y) == '#' {
				emptyRow = false
				FindGalaxy(x, y, galaxies).delta.x = expand_x
			}
		}
		if emptyRow {
			expand_x++
		}
	}

	result := 0
	for i := range galaxies {
		a := Point{
			galaxies[i].orig.x + galaxies[i].delta.x,
			galaxies[i].orig.y + galaxies[i].delta.y,
		}
		for j := range galaxies {
			b := Point{
				galaxies[j].orig.x + galaxies[j].delta.x,
				galaxies[j].orig.y + galaxies[j].delta.y,
			}
			result += abs(a.x-b.x) + abs(a.y-b.y)
		}
	}

	fmt.Println(result / 2)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
