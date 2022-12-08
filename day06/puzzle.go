package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	signal := scanner.Text()

	fmt.Println("first part:", startOfPacketIndex(signal))
	fmt.Println("second part:", startOfMessageIndex(signal))
}

func startOfMessageIndex(signal string) int {
	length := len(signal)
	tail := 0

	for head := 14; head < length; head += 1 {
		chunk := signal[tail:head]

		if areDistinct(chunk) {
			return head
		}

		tail += 1
	}

	return 0
}

func startOfPacketIndex(signal string) int {
	length := len(signal)
	tail := 0

	for head := 4; head < length; head += 1 {
		chunk := signal[tail:head]

		if areDistinct(chunk) {
			return head
		}

		tail += 1
	}

	return 0
}

func areDistinct(chunk string) bool {
	for _, char := range chunk {
		if strings.Count(chunk, string(char)) > 1 {
			return false
		}
	}

	return true
}
