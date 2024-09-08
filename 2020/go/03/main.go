package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type XY struct {
	x int
	y int
}

func get_count(slope XY, treeMap [][]bool) int {
	pos := XY{
		x: 0,
		y: 0,
	}
	count := 0
	for ; pos.y < len(treeMap); pos.x, pos.y = (pos.x+slope.x)%len(treeMap[0]), pos.y+slope.y {
		if treeMap[pos.y][pos.x] {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open("03/input.txt")
	// file, err := os.Open("03/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var treeMap [][]bool

	for scanner.Scan() { // internally, it advances token based on sperator
		var line []bool
		for _, char := range scanner.Bytes() {
			switch char {
			case '.':
				line = append(line, false)
				break
			case '#':
				line = append(line, true)
				break
			default:
				fmt.Printf("bad char: %c", char)
			}
		}
		treeMap = append(treeMap, line)
	}

	// part 1
	fmt.Println(get_count(XY{x: 3, y: 1}, treeMap))

	// part 2
	slopes := []XY{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}
	result := 1
	for _, slope := range slopes {
		result *= get_count(slope, treeMap)
	}
	fmt.Println(result)
}
