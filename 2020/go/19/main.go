package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var RULE11DEPTH = 10 // increase if part2 is incorrect

func main() {
	file, err := os.Open("19/input.txt")
	// file, err := os.Open("19/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules := make(map[string]string)
	for scanner.Scan() && scanner.Text() != "" {
		parts := strings.Split(scanner.Text(), ": ")
		if parts[1][0] == '"' {
			parts[1] = string(parts[1][1])
		}
		rules[parts[0]] = parts[1]
	}

	messages := []string{}
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	p1match := strings.ReplaceAll(computeRule(rules["0"], rules, false), " ", "")
	p2match := strings.ReplaceAll(computeRule(rules["0"], rules, true), " ", "")
	count1 := 0
	count2 := 0

	for _, message := range messages {
		match1, _ := regexp.MatchString("\\b"+p1match+"\\b", message)
		if match1 {
			count1++
		}

		match2, _ := regexp.MatchString("\\b"+p2match+"\\b", message)
		if match2 {
			count2++
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)

}

func computeRule(rule string, rules map[string]string, part2 bool) string {
	result := "("
	for _, part := range strings.Split(rule, " ") {
		result += " "
		if part == "|" || part == "a" || part == "b" {
			result += part
			continue
		}

		// 8: 42 * n
		if part2 && part == "8" {
			result += "( " + computeRule(rules["42"], rules, part2) + " ) +"
			continue
		}

		// 11: 42 * n + 31 * n
		if part2 && part == "11" {
			rule42 := computeRule(rules["42"], rules, part2)
			rule31 := computeRule(rules["31"], rules, part2)
			result += "( "
			for i := 1; i < RULE11DEPTH; i++ {
				if i != 1 {
					result += " | "
				}
				result += "( " + strings.Repeat(rule42, i) + strings.Repeat(rule31, i) + " )"
			}
			result += " )"
			continue
		}

		result += computeRule(rules[part], rules, part2)
	}
	return result + " )"
}
