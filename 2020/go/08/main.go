package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  int
	executed  bool
}

func main() {
	file, err := os.Open("08/input.txt")
	// file, err := os.Open("08/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := []Instruction{}

	// read instructions
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		operation := parts[0]
		argument, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, Instruction{operation: operation, argument: argument})
	}

	p1, _ := runProgram(instructions)
	fmt.Println(p1)

	// reset instructions
	for i := range instructions {
		instructions[i].executed = false
	}

	var p2 int
	for i, instruction := range instructions {
		var newOperation string
		switch instruction.operation {
		case "jmp":
			newOperation = "nop"
		case "nop":
			newOperation = "jmp"
		default:
			continue
		}
		newProgram := make([]Instruction, len(instructions))
		copy(newProgram, instructions)
		newProgram[i] = Instruction{operation: newOperation, argument: instruction.argument}
		acc, ok := runProgram(newProgram)
		if !ok {
			continue
		}
		p2 = acc
		break
	}
	fmt.Println(p2)
}

func runProgram(instructions []Instruction) (int, bool) {
	currentLine := 0
	acc := 0
	for true {
		if currentLine == len(instructions) {
			break
		}
		instruction := &instructions[currentLine]
		if instruction.executed {
			return acc, false
		}
		switch instruction.operation {
		case "acc":
			acc += instruction.argument
		case "jmp":
			currentLine += instruction.argument
			instruction.executed = true
			continue
		}
		instruction.executed = true
		currentLine++
	}
	return acc, true
}
