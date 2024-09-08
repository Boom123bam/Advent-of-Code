package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func findSeatID(line string) int {
	rMin, rMax := 0, 127
	cMin, cMax := 0, 7
	for i := 0; i < len(line); i++ {
		char := line[i]
		switch char {
		case 'F':
			rMax = (rMax + rMin) / 2
		case 'B':
			rMin = (rMax+rMin)/2 + 1
		case 'L':
			cMax = (cMax + cMin) / 2
		case 'R':
			cMin = (cMax+cMin)/2 + 1
		}
	}
	return rMin*8 + cMin
}

func main() {
	file, err := os.Open("05/input.txt")
	// file, err := os.Open("05/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// part 1
	maxID := 0
	var IDs []int

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		ID := findSeatID(scanner.Text())
		IDs = append(IDs, ID)
		maxID = max(maxID, ID)
	}
	fmt.Println(maxID)

	// part 2
	sort.Ints(IDs)
	var prevID int
	for i, ID := range IDs {
		if i != 0 && ID != prevID+1 {
			fmt.Println(ID - 1)
		}
		prevID = ID
	}
}
