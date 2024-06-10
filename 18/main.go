package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("18/input.txt")
	// file, err := os.Open("18/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		sum += eval(scanner.Text())
	}
	fmt.Println(sum)
}

func eval(input string) int {
	result := 0
	operator := byte('+')
	for i := 0; i < len(input); i++ {
		char := input[i]
		switch char {
		case ' ':
		case '+':
			operator = char
		case '*':
			operator = char
		default:
			var operand int
			if char == '(' {
				start := i
				i = getEndBracket(input, i)
				operand = eval(input[start+1 : i])
			} else {
				operand = int(char - '0')
			}
			if operator == '+' {
				result += operand
			} else {
				result *= operand
			}
		}
	}
	return result
}

func getEndBracket(input string, startBracketIndex int) int {
	nestCount := 1
	i := startBracketIndex + 1
	for ; i < len(input); i++ {
		switch input[i] {
		case '(':
			nestCount++
		case ')':
			nestCount--
		}
		if nestCount == 0 {
			return i
		}
	}
	return -1
}
