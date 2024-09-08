package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Instruction interface {
	instruction()
}

type SetMemInstruction struct {
	address uint64
	value   uint64
}

type SetMaskInstruction struct {
	mask string
}

func (s SetMemInstruction) instruction()  {}
func (s SetMaskInstruction) instruction() {}

func main() {
	file, err := os.Open("14/input.txt")
	// file, err := os.Open("14/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := []Instruction{}

	for scanner.Scan() {
		line := scanner.Text()

		if line[:4] == "mask" {
			instructions = append(instructions, SetMaskInstruction{mask: line[7:]})
		} else {
			nums := extract_nums(line)

			instruction := SetMemInstruction{
				address: uint64(nums[0]),
				value:   uint64(nums[1]),
			}

			instructions = append(instructions, instruction)
		}
	}

	// part 1
	mask := ""
	memory := make(map[uint64]uint64)
	for _, instruction := range instructions {
		switch instruction := instruction.(type) {
		case SetMaskInstruction:
			mask = instruction.mask
		case SetMemInstruction:
			memory[instruction.address] = apply_mask(mask, instruction.value)
		}
	}

	var sum uint64
	for _, val := range memory {
		sum += val
	}
	fmt.Println(sum)

	// part 2
	memory = make(map[uint64]uint64)
	for _, instruction := range instructions {
		switch instruction := instruction.(type) {
		case SetMaskInstruction:
			mask = instruction.mask
		case SetMemInstruction:
			addresses := get_all_addresses(mask, instruction.address)
			for _, address := range addresses {
				memory[address] = instruction.value
			}
		}
	}

	sum = 0
	for _, val := range memory {
		sum += val
	}
	fmt.Println(sum)

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

func apply_mask(mask string, num uint64) uint64 {
	for i, j := 0, len(mask)-1; i < len(mask); i, j = i+1, j-1 {
		switch mask[j] {
		case '1':
			num = setBit(num, i)
		case '0':
			num = clearBit(num, i)
		}
	}
	return num
}

func get_all_addresses(mask string, address uint64) []uint64 {
	floating_bits := []int{}
	for i, j := 0, len(mask)-1; i < len(mask); i, j = i+1, j-1 {
		switch mask[j] {
		case '1':
			address = setBit(address, i)
		case 'X':
			address = clearBit(address, i)
			floating_bits = append(floating_bits, i)
		}
	}
	addresses := []uint64{}
	for i := 0; i < int(math.Pow(float64(2), float64(len(floating_bits)))); i++ {
		i_in_binary := to_bits(i, len(floating_bits))
		new_address := address
		for i, bit := range i_in_binary {
			if bit {
				new_address = setBit(new_address, floating_bits[i])
			}
		}
		addresses = append(addresses, new_address)
	}
	return addresses
}

func to_bits(num, length int) []bool {
	result := make([]bool, length)
	for i, exp := 0, 1; i < length; i, exp = i+1, exp*2 {
		if num&exp != 0 {
			result[i] = true
		}
	}
	return result
}

func setBit(n uint64, pos int) uint64 {
	n |= (1 << pos)
	return n
}
func clearBit(n uint64, pos int) uint64 {
	mask := uint64(^(1 << pos))
	n &= mask
	return n
}
