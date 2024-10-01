package main

import "fmt"

func main() {
	input := "712643589"
	// input := "389125467"

	cups := make([]int, 10)
	first := int(input[0] - '0')
	for i := 0; i < 8; i++ {
		current := int(input[i] - '0')
		next := int(input[i+1] - '0')
		cups[current] = next
	}
	cups[int(input[8]-'0')] = first

	cups1 := make([]int, 10)
	copy(cups1, cups)

	cups1 = run(cups1, 10, first)

	c := 1
	for i := 0; i < 8; i++ {
		fmt.Print(cups1[c])
		c = cups1[c]
	}
	fmt.Println()

	// part 2

	cups[int(input[8]-'0')] = 10
	for i := 10; i < 1000000; i++ {
		cups = append(cups, i+1)
	}
	cups = append(cups, first)

	cups = run(cups, 10000000, first)
	fmt.Println(cups[1] * cups[cups[1]])
}

func run(cups []int, times int, first int) []int {
	c := first
	for i := 0; i < times; i++ {
		x := c
		m := []int{}
		for j := 0; j < 3; j++ {
			x = cups[x]
			m = append(m, x)

		}
		d := c - 1
		if d == 0 {
			d = len(cups) - 1
		}

		for in(m, d) {
			d = d - 1
			if d == 0 {
				d = len(cups) - 1
			}
		}
		n := cups[x]
		cups[x] = cups[d]
		cups[d] = cups[c]
		cups[c] = n
		c = n
	}
	return cups

}

func printCups(start, highlight int, cups []int) {
	current := start
	for cups[current] != start {
		if current == highlight {
			fmt.Printf("(%d)", current)
		} else {
			fmt.Printf(" %d ", current)
		}
		current = cups[current]
	}
	if current == highlight {
		fmt.Printf("(%d)", current)
	} else {
		fmt.Printf(" %d ", current)
	}
	fmt.Println()
}

func in(nums []int, target int) bool {
	for _, n := range nums {
		if n == target {
			return true
		}
	}
	return false
}
