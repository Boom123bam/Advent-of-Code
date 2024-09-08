package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grid [][]rune

func main() {
	file, err := os.Open("11/input.txt")
	// file, err := os.Open("11/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	original := Grid{}

	row := 0
	for scanner.Scan() {
		original = append(original, []rune{})
		for _, char := range scanner.Text() {
			original[row] = append(original[row], char)
		}
		row++
	}

	// part1
	grid := original.Copy()
	newGrid := grid.Copy().Permutate()
	for !newGrid.Equals(grid) {
		grid = newGrid.Copy()
		newGrid = newGrid.Permutate()
	}
	fmt.Println(newGrid.CountOccupiedSeats())

	// part2
	grid = original.Copy()
	newGrid = grid.Copy().Permutate2()
	for !newGrid.Equals(grid) {
		grid = newGrid.Copy()
		newGrid = newGrid.Permutate2()
	}

	fmt.Println(newGrid.CountOccupiedSeats())

}

func (grid Grid) String() string {
	sb := strings.Builder{}
	for _, line := range grid {
		for _, char := range line {
			sb.WriteRune(char)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (grid Grid) Copy() Grid {
	duplicate := make(Grid, len(grid))
	for i := range grid {
		duplicate[i] = make([]rune, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}

func (grid Grid) Permutate() Grid {
	copy := grid.Copy()
	for r, row := range grid {
		for c := range row {
			switch grid[r][c] {
			case '.':
				continue
			case 'L':
				if grid.CountAdjacentOccupiedSeats(r, c) == 0 {
					copy[r][c] = '#'
				}
			case '#':
				if grid.CountAdjacentOccupiedSeats(r, c) >= 4 {
					copy[r][c] = 'L'
				}
			}
		}
	}
	return copy
}

func (grid Grid) Permutate2() Grid {
	copy := grid.Copy()
	for r, row := range grid {
		for c := range row {
			switch grid[r][c] {
			case '.':
				continue
			case 'L':
				if grid.CountVisibleOccupiedSeats(r, c) == 0 {
					copy[r][c] = '#'
				}
			case '#':
				if grid.CountVisibleOccupiedSeats(r, c) >= 5 {
					copy[r][c] = 'L'
				}
			}
		}
	}
	return copy
}

func (grid Grid) CountAdjacentOccupiedSeats(row, col int) int {
	result := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if !grid.isValidIndex(r, c) {
				continue
			}
			if r == row && c == col {
				continue
			}
			if grid[r][c] == '#' {
				result++
			}
		}
	}
	return result
}

func (grid Grid) CountVisibleOccupiedSeats(row, col int) int {
	result := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			multiplier := 1
			for grid.isValidIndex(row+dy*multiplier, col+dx*multiplier) {
				if grid[row+dy*multiplier][col+dx*multiplier] == '#' {
					result++
					break
				}
				if grid[row+dy*multiplier][col+dx*multiplier] == 'L' {
					break
				}
				multiplier++
			}
		}
	}
	return result
}

func (grid Grid) Equals(grid2 Grid) bool {
	if (len(grid) != len(grid2)) || (len(grid[0]) != len(grid2[0])) {
		return false
	}
	for r, row := range grid {
		for c := range row {
			if grid[r][c] != grid2[r][c] {
				return false
			}
		}
	}
	return true
}

func (grid Grid) CountOccupiedSeats() int {
	count := 0
	for r, row := range grid {
		for c := range row {
			if grid[r][c] == '#' {
				count++
			}
		}
	}
	return count
}

func (grid Grid) isValidIndex(row, col int) bool {
	if row < 0 || row >= len(grid) {
		return false
	}
	if col < 0 || col >= len(grid[0]) {
		return false
	}
	return true
}
