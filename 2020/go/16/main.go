package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FieldData struct {
	name    string
	a_start int
	a_end   int
	b_start int
	b_end   int
}

type Ticket []int

func main() {
	file, err := os.Open("16/input.txt")
	// file, err := os.Open("16/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fieldDatas := []FieldData{}

	scanner.Scan()
	row := scanner.Text()
	for row != "" {
		parts := strings.Split(row, ":")
		nums := extract_nums(parts[1])

		fieldDatas = append(fieldDatas, FieldData{
			name:    parts[0],
			a_start: nums[0],
			a_end:   nums[1],
			b_start: nums[2],
			b_end:   nums[3],
		})
		scanner.Scan()
		row = scanner.Text()
	}
	scanner.Scan()
	scanner.Scan()
	my_ticket := extract_nums(scanner.Text())
	scanner.Scan()
	scanner.Scan()
	nearby_tickets := []Ticket{}
	for scanner.Scan() {
		nearby_tickets = append(nearby_tickets, extract_nums(scanner.Text()))
	}

	error_rate := 0
	for i := 0; i < len(nearby_tickets); {
		ticket := nearby_tickets[i]
		invalid := false
		for _, val := range ticket {
			if !has_valid_field(fieldDatas, val) {
				error_rate += val
				invalid = true
			}
		}
		if invalid {
			nearby_tickets = append(nearby_tickets[:i], nearby_tickets[i+1:]...)
			continue
		}
		i++
	}
	fmt.Println(error_rate)

	all_tickets := append(nearby_tickets, my_ticket)

	// possible_positions_lengths := make([]int, len(fieldDatas))
	possible_positions := make([][]int, len(fieldDatas))
	// for each field find which positions are possible
	for f, fieldData := range fieldDatas {
		for i := 0; i < len(all_tickets[0]); i++ {
			if fieldData.is_possible_pos(i, all_tickets) {
				possible_positions[f] = append(possible_positions[f], i)
			}
		}
	}
	eliminate(&possible_positions, []int{})

	final_positions := make([]FieldData, len(fieldDatas))

	for i, fieldData := range fieldDatas {
		final_positions[possible_positions[i][0]] = fieldData
	}

	ans := 1
	for i, fieldData := range final_positions {
		if len(fieldData.name) >= 9 && fieldData.name[:9] == "departure" {
			ans *= my_ticket[i]
		}
	}
	fmt.Println(ans)
}

func eliminate(possible_positions *[][]int, known []int) {
	for i, poss := range *possible_positions {
		if has(known, i) {
			continue
		}
		if len(poss) == 1 {
			known = append(known, i)
			num_to_remove := poss[0]
			for j := range *possible_positions {
				if j == i {
					continue
				}
				(*possible_positions)[j] = remove_first_occurance((*possible_positions)[j], num_to_remove)
			}
			eliminate(possible_positions, known)
			return
		}
	}
}

func remove_first_occurance(nums []int, target int) []int {
	for i, num := range nums {
		if num == target {
			return append(nums[:i], nums[i+1:]...)
		}
	}
	return nums
}

func has(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func (field FieldData) is_possible_pos(i int, tickets []Ticket) bool {
	for _, ticket := range tickets {
		if !field.in_either_range(ticket[i]) {
			return false
		}
	}
	return true
}

func has_valid_field(fieldDatas []FieldData, num int) bool {
	for _, fieldData := range fieldDatas {
		if fieldData.in_either_range(num) {
			return true
		}
	}
	return false
}

func (field FieldData) in_either_range(num int) bool {
	if num >= field.a_start && num <= field.a_end {
		return true
	}
	if num >= field.b_start && num <= field.b_end {
		return true
	}
	return false
}

func extract_nums(line string) []int {
	result := []int{}
	sum := 0
	prev_num := false
	for _, char := range line {
		num, err := strconv.Atoi(string(char))
		if err == nil {
			sum = sum*10 + num
			prev_num = true
			continue
		}
		if prev_num {
			result = append(result, sum)
		}
		prev_num = false
		sum = 0
	}
	if prev_num {
		result = append(result, sum)
	}
	return result
}
