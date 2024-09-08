package main

import "fmt"

func main() {
	input := []int{12, 20, 0, 6, 1, 17, 7}
	fmt.Println(solve(input, 2020))
	fmt.Println(solve(input, 30000000))
}

func solve(nums []int, max int) int {
	spoken := make(map[int][]int)

	for i, element := range nums {
		spoken[element] = []int{i + 1}
	}

	lastSpoken := nums[len(nums)-1]
	for i := len(nums); i < max; i++ {
		curr, total := spoken[lastSpoken], 0

		if len(curr) <= 1 {
			spoken[0] = append(spoken[0], i+1)
		} else {
			total = abs(curr[len(curr)-1] - curr[len(curr)-2])
			spoken[total] = append(spoken[total], i+1)
		}

		lastSpoken = total
	}

	return lastSpoken
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
