package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Boom123bam/advent-of-go-2023/utils"
)

var nums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func wordToNum(word string) int {
	for n, num := range nums {
		if word == num {
			return n + 1
		}
	}
	return -1
}

func findFirstNum(line string, includeLetters bool) int {
	for i := 0; i < len(line); i++ {
		if utils.ToNum(line[i]) != -1 {
			return utils.ToNum(line[i])
		}
		if !includeLetters {
			continue
		}
		for ahead := 3; ahead <= 5; ahead++ {
			if i+ahead > len(line) {
				break
			}
			if wordToNum(line[i:i+ahead]) != -1 {
				return wordToNum(line[i : i+ahead])
			}
		}
	}
	return -1
}

func findLastNum(line string, includeLetters bool) int {
	for i := len(line) - 1; i >= 0; i-- {
		if utils.ToNum(line[i]) != -1 {
			return utils.ToNum(line[i])
		}
		if !includeLetters {
			continue
		}
		for ahead := 3; ahead <= 5; ahead++ {
			if i-ahead < 0 {
				break
			}
			if wordToNum(line[i+1-ahead:i+1]) != -1 {
				return wordToNum(line[i+1-ahead : i+1])
			}
		}
	}
	return -1
}

func main() {
	file, err := os.Open("01/input.txt")
	// file, err := os.Open("01/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0
	for scanner.Scan() {
		part1 += findFirstNum(scanner.Text(), false)*10 + findLastNum(scanner.Text(), false)
		part2 += findFirstNum(scanner.Text(), true)*10 + findLastNum(scanner.Text(), true)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
