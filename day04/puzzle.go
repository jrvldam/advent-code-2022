package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	lines := getLines(file)

	fmt.Println("first part:", findAllFullyOverlappedAssignments(lines))
	fmt.Println("second part:", findAllOverlappedAssignments(lines))
}

func findAllOverlappedAssignments(lines [][]string) int {
	var totalOverlaps int
	for _, line := range lines {
		if hasOverlapped(line) {
			totalOverlaps += 1
		}
	}

	return totalOverlaps
}

func hasOverlapped(assignments []string) bool {
	left, right := getRangesFromAssignments(assignments)

	if len(left) >= len(right) {
		for _, n := range right {
			if isAInAs(n, left) {
				return true
			}
		}
	} else {
		for _, n := range left {
			if isAInAs(n, right) {
				return true
			}
		}
	}

	return false
}

func findAllFullyOverlappedAssignments(lines [][]string) int {
	var totalOverlaps int
	for _, line := range lines {
		if hasFullOverlapped(line) {
			totalOverlaps += 1
		}
	}

	return totalOverlaps
}

func hasFullOverlapped(assignments []string) bool {
  left, right := getRangesFromAssignments(assignments)

	if len(left) >= len(right) {
		for _, n := range right {
			if !isAInAs(n, left) {
				return false
			}
		}
	} else {
		for _, n := range left {
			if !isAInAs(n, right) {
				return false
			}
		}
	}

	return true
}

func isAInAs(a int, as []int) bool {
	for _, n := range as {
		if n == a {
			return true
		}
	}

	return false
}

func getRangesFromAssignments(assignments []string) (left []int, right []int) {
	leftLower, leftUpper := getBounds(assignments[0])
	rightLower, rightUpper := getBounds(assignments[1])

	left = makeRange(leftLower, leftUpper)
	right = makeRange(rightLower, rightUpper)

	return left, right
}

func makeRange(lower, upper int) []int {
	var series []int
	for i := lower; i <= upper; i += 1 {
		series = append(series, i)
	}

	return series
}

func getBounds(side string) (int, int) {
	bounds := strings.Split(side, "-")
	lower, _ := strconv.ParseInt(bounds[0], 10, 32)
	upper, _ := strconv.ParseInt(bounds[1], 10, 32)

	return int(lower), int(upper)
}

func getLines(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)

	var lines [][]string
	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		lines = append(lines, assignments)
	}
	return lines
}
