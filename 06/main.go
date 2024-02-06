package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("06/input.txt")
	// file, err := os.Open("06/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	p1count := 0
	p2count := 0

	scanner.Scan()
	line := scanner.Text()
	allLetters := getLetters(line)
	commonLetters := getLetters(line)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// add count
			p1count += len(allLetters)
			p2count += len(commonLetters)

			// get the next group
			if !scanner.Scan() {
				break
			}
			line = scanner.Text()
			allLetters = getLetters(line)
			commonLetters = getLetters(line)
		}
		addLetters(line, allLetters)
		filterLetters(line, commonLetters)
	}
	fmt.Println(p1count)
	fmt.Println(p2count)
}

func getLetters(line string) map[byte]bool {
	letters := map[byte]bool{}
	for i := 0; i < len(line); i++ {
		letter := line[i]
		letters[letter] = true
	}
	return letters
}

func filterLetters(line string, letters map[byte]bool) {
	// delete letters in map that are not in line
	for letter := range letters {
		if !strings.Contains(line, string(letter)) {
			delete(letters, letter)
		}
	}
}

func addLetters(line string, letters map[byte]bool) {
	// add all letters in line to map
	for i := 0; i < len(line); i++ {
		letter := line[i]
		letters[letter] = true
	}
}
