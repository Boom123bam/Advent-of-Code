package main

import "fmt"

func main() {
	cardPubKey := 18499292
	doorPubKey := 8790390
	// cardPubKey := 5764801
	// doorPubKey := 17807724

	cardLoopSize := -1
	doorLoopSize := -1

	val := 1
	subjectNum := 7
	for i := 0; cardLoopSize == -1 || doorLoopSize == -1; i++ {
		val *= subjectNum
		val %= 20201227
		if cardLoopSize == -1 && val == cardPubKey {
			cardLoopSize = i + 1
		}
		if doorLoopSize == -1 && val == doorPubKey {
			doorLoopSize = i + 1
		}
	}

	val = 1
	subjectNum = doorPubKey
	for i := 0; i < cardLoopSize; i++ {
		val *= subjectNum
		val %= 20201227
	}
	fmt.Println(val)
}
