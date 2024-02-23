package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagCol string
type bagList map[bagCol](map[bagCol]int)

func main() {
	file, err := os.Open("07/input.txt")
	// file, err := os.Open("07/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	bags := bagList{}

	// read bags
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		i := 0
		for line[i] != "contain" {
			i++
		}
		parentCol := bagCol(strings.Join(line[:i-1], " "))
		bags[parentCol] = make(map[bagCol]int)
		i++
		for i < len(line) {
			quantity, err := strconv.Atoi(line[i])
			if err != nil {
				break
			}
			i++
			lineNameStart := i
			for line[i][:3] != "bag" {
				i++
			}
			childCol := bagCol(strings.Join(line[lineNameStart:i], " "))
			bags[parentCol][childCol] = quantity
			i++
		}
	}
	count := 0
	p2 := 0
	for col := range bags {
		if col == "shiny gold" {
			p2 = countBags(col, bags)
		}
		if hasShinyGoldBag(col, bags) {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(p2)
}

func hasShinyGoldBag(parentCol bagCol, bags bagList) bool {
	bagData := bags[parentCol]
	for col := range bagData {
		if col == "shiny gold" || hasShinyGoldBag(col, bags) {
			return true
		}
	}
	return false
}

func countBags(parentCol bagCol, bags bagList) int {
	total := 0
	bagData := bags[parentCol]
	for col := range bagData {
		total += bagData[col] * (countBags(col, bags) + 1)
	}
	return total
}
