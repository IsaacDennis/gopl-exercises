package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Println("1. Type something: ")
	data, err := input.ReadBytes('\n')
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	sha1 := sha256.Sum256(data)
	fmt.Println("2. Type something: ")
	data, err = input.ReadBytes('\n')
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	sha2 := sha256.Sum256(data)
	fmt.Printf("Bit difference: %d\n", DiffSHA256(sha1, sha2))
}

func DiffSHA256(sha1, sha2 [32]byte) int {
	var diff int
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			diff += int(((sha1[i] >> j) & 1) ^ ((sha2[i] >> j) & 1))
		}
	}
	return diff
}
