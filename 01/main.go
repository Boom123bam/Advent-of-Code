package main

import (
	"fmt"
	"io"
	"os"
)

func find2(numbers []int) int {
	for _, num1 := range numbers {
		for _, num2 := range numbers {
			if num1+num2 == 2020 {
				return num1 * num2
			}
		}
	}
	return -1
}

func find3(numbers []int) int {
	for _, num1 := range numbers {
		for _, num2 := range numbers {
			for _, num3 := range numbers {
				if num1+num2+num3 == 2020 {
					return num1 * num2 * num3
				}
			}
		}
	}
	return -1
}

func main() {
	file, error := os.Open("01/input.txt")
	// file, error := os.Open("01/test.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
	var numbers []int
	currentNum := 0
	for _, ascii := range buf {
		switch ascii {
		case 10:
			numbers = append(numbers, currentNum)
			currentNum = 0
			break
		case 0:
			break
		default:
			currentNum = currentNum*10 + int(ascii) - 48
		}
	}
	fmt.Println(find2(numbers))
	fmt.Println(find3(numbers))

}
