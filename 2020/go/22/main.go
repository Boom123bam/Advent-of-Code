package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("22/input.txt")
	// file, err := os.Open("22/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	p1 := []int{}
	p2 := []int{}
	scanner.Scan()
	for scanner.Scan() && scanner.Text() != "" {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		p1 = append(p1, num)
	}

	scanner.Scan()
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		p2 = append(p2, num)
	}

	p1copy := make([]int, len(p1))
	p2copy := make([]int, len(p2))
	copy(p1copy, p1)
	copy(p2copy, p2)

	for len(p1copy) > 0 && len(p2copy) > 0 {
		p1copy, p2copy = playRound1(p1copy, p2copy)
	}

	if len(p1copy) == 0 {
		fmt.Println(getScore(p2copy))
	} else {
		fmt.Println(getScore(p1copy))
	}

	_, deck := playGame(p1, p2)
	fmt.Println(getScore(deck))
}

func playRound1(p1, p2 []int) ([]int, []int) {
	p1card := p1[0]
	p2card := p2[0]
	p1 = p1[1:]
	p2 = p2[1:]

	if p1card > p2card {
		p1 = append(p1, p1card)
		p1 = append(p1, p2card)
	} else {
		p2 = append(p2, p2card)
		p2 = append(p2, p1card)
	}
	return p1, p2
}

func playGame(p1, p2 []int) (int, []int) {
	seenDecks := []string{}
	for true {
		winner := 0
		currentDeck := intSliceToString(p1) + "|" + intSliceToString(p2)
		for _, deck := range seenDecks {
			if currentDeck == deck {
				return 1, p1
			}
		}
		seenDecks = append(seenDecks, currentDeck)

		// play round
		p1card := p1[0]
		p2card := p2[0]
		p1 = p1[1:]
		p2 = p2[1:]

		// both have at least as many cards as the value of the card drawn
		if len(p1) >= p1card && len(p2) >= p2card {
			p1copy := make([]int, len(p1))
			p2copy := make([]int, len(p2))
			copy(p1copy, p1)
			copy(p2copy, p2)
			winner, _ = playGame(p1copy[:p1card], p2copy[:p2card])
		} else {
			if p1card > p2card {
				winner = 1
			} else {
				winner = 2
			}
		}

		// higher value wins round
		if winner == 1 {
			p1 = append(p1, p1card)
			p1 = append(p1, p2card)
		} else {
			p2 = append(p2, p2card)
			p2 = append(p2, p1card)
		}
		if len(p1) == 0 {
			return 2, p2
		} else if len(p2) == 0 {
			return 1, p1
		}
	}
	return -1, []int{}
}

func intSliceToString(intSlice []int) string {
	stringSlice := make([]string, len(intSlice))
	for i, num := range intSlice {
		stringSlice[i] = strconv.Itoa(num)
	}
	return strings.Join(stringSlice, ",")
}

func getScore(deck []int) int {
	result := 0
	for i, j := 0, len(deck); i < len(deck); i, j = i+1, j-1 {
		result += deck[i] * j
	}
	return result
}
