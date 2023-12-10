package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

type Point struct{ x, y int }

type Grid struct {
	width int
	tiles []byte
}

func NewGrid(filename string) Grid {
	tiles, _ := os.ReadFile(filename)
	width := bytes.IndexByte(tiles, '\n') + 1
	return Grid{width, tiles}
}

func (g *Grid) At(x, y int) byte {
	return g.tiles[x+y*g.width]
}

func (g *Grid) PosFromIndex(i int) Point {
	return Point{i % g.width, i / g.width}
}

func (g *Grid) NearPipes(p Point) (Point, Point) {
	pipes := [2]Point{}
	i := 0

	if tile := g.At(p.x, p.y-1); tile == '7' || tile == 'F' || tile == '|' {
		pipes[i] = Point{p.x, p.y - 1}
		i++
	}
	if tile := g.At(p.x, p.y+1); tile == 'L' || tile == 'J' || tile == '|' {
		pipes[i] = Point{p.x, p.y + 1}
		i++
	}
	if tile := g.At(p.x-1, p.y); tile == 'L' || tile == 'F' || tile == '-' {
		pipes[i] = Point{p.x - 1, p.y}
		i++
	}
	if tile := g.At(p.x+1, p.y); tile == '7' || tile == 'J' || tile == '-' {
		pipes[i] = Point{p.x + 1, p.y}
		i++
	}
	return pipes[0], pipes[1]
}

func (g *Grid) Next(p, before Point) Point {
	var next Point
	switch g.At(p.x, p.y) {
	case '-':
		next = Point{
			p.x + (p.x - before.x),
			p.y,
		}
	case '|':
		next = Point{
			p.x,
			p.y + (p.y - before.y),
		}
	case 'L':
		next = Point{
			p.x + (p.y - before.y),
			p.y + (p.x - before.x),
		}
	case 'F':
		next = Point{
			p.x - (p.y - before.y),
			p.y - (p.x - before.x),
		}
	case '7':
		next = Point{
			p.x + (p.y - before.y),
			p.y + (p.x - before.x),
		}
	case 'J':
		next = Point{
			p.x - (p.y - before.y),
			p.y - (p.x - before.x),
		}
	}
	return next
}

func main() {
	filename := flag.String("i", "input", "input file")
	flag.Parse()
	grid := NewGrid(*filename)

	startTile := grid.PosFromIndex(bytes.IndexByte(grid.tiles, 'S'))
	pipe1, _ := grid.NearPipes(startTile)

	i := 0
	for tile, before := pipe1, startTile; tile != startTile; i++ {
		tile, before = grid.Next(tile, before), tile
	}

	fmt.Println(i/2 + 1)
}
