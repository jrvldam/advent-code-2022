package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	bags := getBags(inputFile)
	fmt.Println("first part:", getPrioritySum(bags))
	fmt.Println("second part:", getBadgesSum(bags))
}

func getBadgesSum(bags [][]string) int {
	var totalBadges int
	for i := 0; i < len(bags); i += 3 {
		group := bags[i : i+3]
		badge := getBadge(group)
		if badge != "" {
			totalBadges += strings.Index(chars, badge) + 1
		}
	}

	return totalBadges
}

func getBadge(group [][]string) string {
	var bags []string
	for i := 0; i < len(group); i += 1 {
		bags = append(bags, group[i][0]+group[i][1])
	}

	badge := findCommonBadge(bags[0], bags[1], bags[2])

	return badge
}

func findCommonBadge(first string, second string, third string) string {
	for i := 0; i < len(first); i += 1 {
		char := first[i : i+1]

		if strings.Contains(second, char) && strings.Contains(third, char) {
			return char
		}
	}

	return ""
}

func getPrioritySum(bags [][]string) int {
	var totalPriority int
	for _, bag := range bags {
		left := bag[0]
		right := bag[1]
		coincidence := findCoincidence(left, right)

		if coincidence != "" {
			priority := strings.Index(chars, coincidence) + 1
			totalPriority += priority
		}
	}

	return totalPriority
}

func findCoincidence(left string, right string) string {
	for i := 0; i < len(left); i += 1 {
		char := left[i : i+1]
		if strings.Contains(right, char) {
			return char
		}
	}

	return ""
}

func getBags(inputFile *os.File) [][]string {
	scanner := bufio.NewScanner(inputFile)

	var bags [][]string
	for scanner.Scan() {
		line := scanner.Text()
		bag := getBag(line)
		bags = append(bags, bag)
	}

	return bags
}

func getBag(line string) []string {
	pivot := len(line) / 2
	return []string{
		line[:pivot],
		line[pivot:],
	}
}
