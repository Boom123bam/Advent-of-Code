package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("13/input.txt")
	// file, err := os.Open("13/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	earliest, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	scanner.Scan()

	nums := []int{}
	for _, item := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(item)
		if err == nil {
			nums = append(nums, num)
		} else {
			nums = append(nums, -1)
		}
	}

	min_waiting_mins := 99999
	min_waiting_id := -1

	for _, id := range nums {
		if id == -1 {
			continue
		}
		waiting_mins := id - earliest%id
		if waiting_mins < min_waiting_mins {
			min_waiting_mins = waiting_mins
			min_waiting_id = id
		}

	}

	fmt.Println(min_waiting_mins * min_waiting_id)
	fmt.Println(find_p2_timestamp(nums))
}

func find_p2_timestamp(ids []int) int {
	as := []int{}
	ms := []int{}
	j := 0
	for i := len(ids) - 1; i >= 0; i, j = i-1, j+1 {
		id := ids[i]
		if id == -1 {
			continue
		}
		as = append([]int{j}, as...)
		ms = append([]int{id}, ms...)
	}
	return crt(as, ms) - j + 1
}

func crt(as []int, ms []int) int {
	// chinese remainder theorm
	cs := make([]int, len(as))
	for i := range ms {
		cs[i] = find_c(ms, i)
	}
	bs := make([]int, len(as))
	M := 1
	for i, m := range ms {
		M *= m
		bs[i] = find_b(cs[i], m)
	}

	result := 0
	for i := range as {
		result += as[i] * bs[i] * cs[i]
	}
	for result > M {
		result -= M
	}
	return result
}

func find_b(c, m int) int {
	result := 1
	for (result*c)%m != 1 {
		result++
	}
	return result
}

func find_c(ms []int, index int) int {
	result := 1
	for i, m := range ms {
		if i != index {
			result *= m
		}
	}
	return result
}
