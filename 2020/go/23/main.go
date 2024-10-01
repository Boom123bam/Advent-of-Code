package main

import "fmt"

type Cup struct {
	label int
	next  *Cup
}

func main() {
	input := "712643589"
	// input := "389125467"
	// input := "1"

	// make linked list
	first := &Cup{label: int(input[0] - '0')}
	current := first
	for i := 1; i < len(input); i++ {
		next := &Cup{label: int(input[i] - '0')}
		current.next = next
		current = next
	}

	for i := 10; i <= 1000000; i++ {
		next := &Cup{label: i}
		current.next = next
		current = next
	}

	current.next = first
	current = first

	for i := 0; i < 10000000; i++ {
		if i%10000 == 0 {
			fmt.Println(i)
		}
		// pick up cups
		pickUpStart := current.next
		current.next = current.next.next.next.next

		// get dest cup
		dest := getDestCup(current, current.label-1)

		// drop off 3 cups after
		pickUpStart.next.next.next = dest.next
		dest.next = pickUpStart

		current = current.next
	}

	for current.label != 1 {
		current = current.next
	}
	for current.next.label != 1 {
		fmt.Print(current.next.label)
		current = current.next
	}
	fmt.Println()

}

func getDestCup(start *Cup, targetLabel int) *Cup {
	current := start.next
	if targetLabel == 0 {
		targetLabel = 9
	}
	for current != start {
		if current.label == targetLabel {
			return current
		}
		current = current.next
	}
	return getDestCup(start, targetLabel-1)
}

func printCups(start *Cup, highlight *Cup) {
	current := start
	for current.next != start {
		if current == highlight {
			fmt.Printf("(%d)", current.label)
		} else {
			fmt.Printf(" %d ", current.label)
		}
		current = current.next
	}
	fmt.Printf("%d \n", current.label)
}
