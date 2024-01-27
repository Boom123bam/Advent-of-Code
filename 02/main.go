package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordData struct {
	password string
	num1     int
	num2     int
	letter   byte
}

func countOccurrences(s string, c byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
		}
	}
	return count
}

func main() {
	file, error := os.Open("02/input.txt")
	// file, error := os.Open("02/test.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()
	buf := make([]byte, 1024*64)
	file.Read(buf)
	var lines []string
	line := ""

	// convert file to lines
	for _, ascii := range buf {
		if ascii == 0 {
			break
		}
		if ascii == 10 {
			lines = append(lines, line)
			line = ""
			continue
		}
		line += string(ascii)
	}

	// convert lines to data
	var passwords []PasswordData
	for _, line := range lines {
		parts := strings.Fields(line)
		nums := strings.Split(parts[0], "-")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		passwords = append(passwords, PasswordData{
			num1:     num1,
			num2:     num2,
			letter:   parts[1][0],
			password: parts[2],
		})
	}

	// part1
	p1Count := 0
	for _, passwordData := range passwords {
		count := countOccurrences(passwordData.password, passwordData.letter)
		if count >= passwordData.num1 && count <= passwordData.num2 {
			p1Count++
		}
	}
	fmt.Println(p1Count)

	// part2
	p2Count := 0
	for _, passwordData := range passwords {
		inPos1 := passwordData.password[passwordData.num1-1] == passwordData.letter
		inPos2 := passwordData.password[passwordData.num2-1] == passwordData.letter
		if (inPos1 || inPos2) && !(inPos1 && inPos2) {
			p2Count++
		}
	}
	fmt.Println(p2Count)
}
