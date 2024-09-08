package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("10/input.txt")
	// file, err := os.Open("10/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// 0 for charging outlet
	numbers := []int{0}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	// sort numbers
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	// add in device adapter
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	diffs := [3]int{}

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		diffs[diff-1]++
	}

	part1 := diffs[0] * diffs[2]
	fmt.Println(part1)

	fmt.Println(getNumPossibleArrangements(numbers))
}

func getNumPossibleArrangements(nums []int) int {
	ways := make([]int, len(nums))
	ways[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0 && nums[i]-nums[j] <= 3; j-- {
			ways[i] += ways[j]
		}
	}
	return ways[len(ways)-1]
}
