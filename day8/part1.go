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
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()

	graph := make(map[string][2]string)
	for scanner.Scan() {
		line := scanner.Text()

		graph[line[:3]] = [2]string{line[7:10], line[12:15]}
	}

	count := 0
	for node := "AAA"; ; {
		// This isn't a range-based loop because we don't
		// need an UTF-8 decoder here.
		for i := 0; i < len(directions); i++ {
			if directions[i] == 'L' {
				node = graph[node][0]
			} else {
				node = graph[node][1]
			}
			count++
		}
		if node == "ZZZ" {
			break
		}
	}
	fmt.Println(count)
}
