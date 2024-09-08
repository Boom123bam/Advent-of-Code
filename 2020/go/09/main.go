package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

const preambleLen = 25

func main() {
	file, err := os.Open("09/input.txt")
	// file, err := os.Open("09/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := []int{}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	part1 := -1

	for i := preambleLen; i < len(numbers); i++ {
		if !hasPossibleSum(numbers[i-preambleLen:i], numbers[i]) {
			part1 = numbers[i]
		}
	}
	fmt.Println(part1)

	contiguous := findContiguousNums(numbers, part1)
	part2 := slices.Min(contiguous) + slices.Max(contiguous)
	fmt.Println(part2)
}

func hasPossibleSum(nums []int, target int) bool {
	for _, a := range nums {
		for _, b := range nums {
			if a+b == target {
				return true
			}
		}
	}
	return false
}

func findContiguousNums(nums []int, target int) []int {
	for i, num := range nums {
		sum := num
		j := i + 1
		for sum < target {
			sum += nums[j]
			j++
		}
		if sum == target {
			return nums[i:j]
		}
	}
	return []int{}
}
