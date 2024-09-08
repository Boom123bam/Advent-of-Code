package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	password := []byte{}
	input := []byte("ffykfhsq")
	for i := 0; len(password) < 8; i++ {
		hash := md5.Sum(append(input, []byte(strconv.Itoa(i))...))
		char := getChar(hash)
		if char != 0 {
			password = append(password, char)
		}
	}
	fmt.Println(string(password))
}

func getChar(hash [16]byte) byte {
	if hash[0] != 0 || hash[1] != 0 || hash[2] > 16 {
		return 0
	}
	if hash[2] < 10 {
		return '0' + hash[2]
	}
	return 'a' + hash[2] - 10
}
