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
	left := strings.Split(assignments[0], "-")
	right := strings.Split(assignments[1], "-")
	leftLower, _ := strconv.ParseInt(left[0], 10, 32)
	leftUpper, _ := strconv.ParseInt(left[1], 10, 32)
	rightLower, _ := strconv.ParseInt(right[0], 10, 32)
	rightUpper, _ := strconv.ParseInt(right[1], 10, 32)

	var leftRange []int
	for i := leftLower; i <= leftUpper; i += 1 {
		leftRange = append(leftRange, int(i))
	}
	var rightRange []int
	for i := rightLower; i <= rightUpper; i += 1 {
		rightRange = append(rightRange, int(i))
	}

	if len(leftRange) >= len(rightRange) {
    for _, n := range rightRange {
      if isAInAs(n, leftRange) {
        return true
      }
    }
	} else {
    for _, n := range leftRange {
      if isAInAs(n, rightRange) {
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
	left := strings.Split(assignments[0], "-")
	right := strings.Split(assignments[1], "-")
	leftLower, _ := strconv.ParseInt(left[0], 10, 32)
	leftUpper, _ := strconv.ParseInt(left[1], 10, 32)
	rightLower, _ := strconv.ParseInt(right[0], 10, 32)
	rightUpper, _ := strconv.ParseInt(right[1], 10, 32)

	var leftRange []int
	for i := leftLower; i <= leftUpper; i += 1 {
		leftRange = append(leftRange, int(i))
	}
	var rightRange []int
	for i := rightLower; i <= rightUpper; i += 1 {
		rightRange = append(rightRange, int(i))
	}

	if len(leftRange) >= len(rightRange) {
    for _, n := range rightRange {
      if !isAInAs(n, leftRange) {
        return false
      }
    }
	} else {
    for _, n := range leftRange {
      if !isAInAs(n, rightRange) {
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

type assignment struct {
  series []int
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
