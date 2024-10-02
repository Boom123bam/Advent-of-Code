package main

import (
	"bufio"
	"fmt"
	"os"
)

// y towards ne, x towards e
type Tile struct {
	x int
	y int
}
type Boundaries struct {
	minX int
	maxX int
	minY int
	maxY int
}

func main() {
	file, err := os.Open("24/input.txt")
	// file, err := os.Open("24/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	blackTiles := make(map[Tile]bool)

	for scanner.Scan() {
		line := scanner.Text() + " "
		tile := Tile{}
		for i := 0; i < len(line)-1; {
			switch line[i : i+2] {
			case "nw":
				tile.x--
				tile.y++
				i += 2
			case "ne":
				tile.y++
				i += 2
			case "sw":
				tile.y--
				i += 2
			case "se":
				tile.x++
				tile.y--
				i += 2
			default:
				switch line[i] {
				case 'e':
					tile.x++
					i += 1
				case 'w':
					tile.x--
					i += 1
				default:
					panic("invalid direction")
				}
			}
		}
		_, ok := blackTiles[tile]
		if !ok {
			blackTiles[tile] = true
		} else {
			delete(blackTiles, tile)
		}
	}
	fmt.Println(len(blackTiles))

	for i := 0; i < 100; i++ {
		flipTiles(blackTiles, getBoundaries(blackTiles))
	}
	fmt.Println(len(blackTiles))
}

func getBoundaries(blackTiles map[Tile]bool) Boundaries {
	boundaries := Boundaries{
		minX: 99999,
		minY: 99999,
		maxX: -99999,
		maxY: -99999,
	}
	for tile := range blackTiles {
		boundaries.minX = min(tile.x, boundaries.minX)
		boundaries.maxX = max(tile.x, boundaries.maxX)
		boundaries.minY = min(tile.y, boundaries.minY)
		boundaries.maxY = max(tile.y, boundaries.maxY)
	}
	return boundaries
}

func flipTiles(blackTiles map[Tile]bool, boundaries Boundaries) {
	toFlip := []Tile{}
	for x := boundaries.minX - 1; x <= boundaries.maxX+1; x++ {
		for y := boundaries.minY - 1; y <= boundaries.maxY+1; y++ {
			tile := Tile{x: x, y: y}
			_, isBlack := blackTiles[tile]
			numBlackNeighbours := getNumBlackNeighbours(blackTiles, tile)

			if (isBlack && (numBlackNeighbours == 0 || numBlackNeighbours > 2)) || (!isBlack && numBlackNeighbours == 2) {
				toFlip = append(toFlip, tile)
			}
		}
	}

	for _, tile := range toFlip {
		_, ok := blackTiles[tile]
		if !ok {
			blackTiles[tile] = true
		} else {
			delete(blackTiles, tile)
		}
	}
}

func getNumBlackNeighbours(blackTiles map[Tile]bool, tile Tile) int {
	count := 0
	for xOff := -1; xOff <= 1; xOff++ {
		for yOff := -1; yOff <= 1; yOff++ {
			if xOff == yOff {
				continue
			}
			t := Tile{x: tile.x + xOff, y: tile.y + yOff}
			_, ok := blackTiles[t]
			if ok {
				count++
			}
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
