package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord [4]int

type Bounds struct {
	xMin, xMax int
	yMin, yMax int
	zMin, zMax int
	wMin, wMax int
}

func main() {
	file, err := os.Open("17/input.txt")
	// file, err := os.Open("17/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	active := []Coord{}

	bounds := Bounds{
		xMin: 0, xMax: 0,
		yMin: 0, yMax: 0,
		zMin: 0, zMax: 0,
		wMin: 0, wMax: 0,
	}

	for y := 0; scanner.Scan(); y++ {
		for x, char := range scanner.Text() {
			if char == '#' {
				c := Coord{x, y, 0, 0}
				active = append(active, c)
				bounds.update(c)
			}
		}
	}

	newActive := active
	for i := 0; i < 6; i++ {
		newActive = cycle(newActive, &bounds, false)
	}
	fmt.Println(len(newActive))

	newActive = active
	for i := 0; i < 6; i++ {
		newActive = cycle(newActive, &bounds, true)
	}
	fmt.Println(len(newActive))
}

func cycle(active []Coord, bounds *Bounds, useW bool) []Coord {
	newActive := []Coord{}
	for x := bounds.xMin - 1; x <= bounds.xMax+1; x++ {
		for y := bounds.yMin - 1; y <= bounds.yMax+1; y++ {
			for z := bounds.zMin - 1; z <= bounds.zMax+1; z++ {
				if useW {
					for w := bounds.wMin - 1; w <= bounds.wMax+1; w++ {
						c := Coord{x, y, z, w}
						neighbours := c.countActiveNeighbours(active)
						isActive := c.in(active)
						if (isActive && (neighbours == 2 || neighbours == 3)) ||
							(!isActive && neighbours == 3) {
							newActive = append(newActive, c)
							bounds.update(c)
						}
					}
				} else {
					c := Coord{x, y, z}
					neighbours := c.countActiveNeighbours(active)
					isActive := c.in(active)
					if (isActive && (neighbours == 2 || neighbours == 3)) ||
						(!isActive && neighbours == 3) {
						newActive = append(newActive, c)
						bounds.update(c)
					}
				}
			}
		}
	}
	return newActive
}

func (c Coord) x() int {
	return c[0]
}
func (c Coord) y() int {
	return c[1]
}
func (c Coord) z() int {
	return c[2]
}
func (c Coord) w() int {
	return c[3]
}

func (c1 Coord) isNeighbour(c2 Coord) bool {
	if c1.x() == c2.x() && c1.y() == c2.y() && c1.z() == c2.z() && c1.w() == c2.w() {
		return false
	} else if abs(c1.x()-c2.x()) > 1 || abs(c1.y()-c2.y()) > 1 || abs(c1.z()-c2.z()) > 1 || abs(c1.w()-c2.w()) > 1 {
		return false
	}
	return true
}

func (c Coord) countActiveNeighbours(active []Coord) int {
	count := 0
	for _, coord := range active {
		if c.isNeighbour(coord) {
			count++
		}
	}
	return count
}

func (c Coord) in(coords []Coord) bool {
	for _, coord := range coords {
		if c.x() == coord.x() && c.y() == coord.y() && c.z() == coord.z() && c.w() == coord.w() {
			return true
		}
	}
	return false
}

func (b *Bounds) update(c Coord) {
	b.xMax = max(b.xMax, c.x())
	b.yMax = max(b.yMax, c.y())
	b.zMax = max(b.zMax, c.z())
	b.wMax = max(b.wMax, c.w())
	b.xMin = min(b.xMin, c.x())
	b.yMin = min(b.yMin, c.y())
	b.zMin = min(b.zMin, c.z())
	b.wMin = min(b.wMin, c.w())
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
