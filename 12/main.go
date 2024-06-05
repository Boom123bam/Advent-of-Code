package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	action byte
	amount int
}

type Ship struct {
	direction int
	x         int
	y         int
}

type Shipv2 struct {
	waypointDx int
	waypointDy int
	x          int
	y          int
}

var directions = []byte{'N', 'E', 'S', 'W'}

func main() {
	file, err := os.Open("12/input.txt")
	// file, err := os.Open("12/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := []Instruction{}

	for scanner.Scan() {
		line := scanner.Text()
		action := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println(err)
			return
		}

		instructions = append(instructions, Instruction{
			action: action,
			amount: amount,
		})
	}

	// part 1
	ship := &Ship{
		direction: 1, // East
		x:         0,
		y:         0,
	}

	for _, instruction := range instructions {
		ship.executeInstruction(instruction)
	}
	fmt.Println(abs(ship.x) + abs(ship.y))

	// part 2
	ship2 := &Shipv2{
		x:          0,
		y:          0,
		waypointDx: 10,
		waypointDy: 1,
	}

	for _, instruction := range instructions {
		ship2.executeInstruction(instruction)
	}
	fmt.Println(abs(ship2.x) + abs(ship2.y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (ship *Ship) executeInstruction(instruction Instruction) {
	switch instruction.action {
	case 'N':
		ship.y += instruction.amount
	case 'S':
		ship.y -= instruction.amount
	case 'E':
		ship.x += instruction.amount
	case 'W':
		ship.x -= instruction.amount
	case 'L':
		ship.direction -= instruction.amount / 90
		if ship.direction <= -1 {
			ship.direction += 4
		}
	case 'R':
		ship.direction += instruction.amount / 90
		for ship.direction >= 4 {
			ship.direction -= 4
		}
	case 'F':
		ship.executeInstruction(Instruction{
			action: directions[ship.direction],
			amount: instruction.amount,
		})
	}
}

func (ship *Shipv2) executeInstruction(instruction Instruction) {
	switch instruction.action {
	case 'N':
		ship.waypointDy += instruction.amount
	case 'S':
		ship.waypointDy -= instruction.amount
	case 'E':
		ship.waypointDx += instruction.amount
	case 'W':
		ship.waypointDx -= instruction.amount
	case 'L':
		for i := 0; i < instruction.amount/90; i++ {
			ship.waypointDx, ship.waypointDy = -ship.waypointDy, ship.waypointDx
		}
	case 'R':
		for i := 0; i < instruction.amount/90; i++ {
			ship.waypointDx, ship.waypointDy = ship.waypointDy, -ship.waypointDx
		}
	case 'F':
		ship.x += ship.waypointDx * instruction.amount
		ship.y += ship.waypointDy * instruction.amount
	}
}
