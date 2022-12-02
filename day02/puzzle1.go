package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
A rock
B paper
C scissors

X rock 1
Y paper 2
Z scissors 3

0 lose
3 draw
6 win
*/
type shapes map[string]string

var you = shapes{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}
var opponent = shapes{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}
var shapeValue = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

func main() {
	inputFile, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
		return
	}

	// result := getRoundsResult(inputFile)
	result := getRoundsResult2(inputFile)

	fmt.Println(result)
}

func getRoundsResult2(inputFile *os.File) int {
	scanner := bufio.NewScanner(inputFile)

	var result int
	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Split(line, " ")
		its := opponent[chunks[0]]
		outcome := getYourOutcome(chunks[1])
		value := getShapeValue(outcome, its)
		result = result + outcome + value
	}

	return result
}

func getShapeValue(outcome int, its string) int {
	switch {
	case outcome == 6:
		if its == "rock" {
			return shapeValue["paper"]
		} else if its == "paper" {
			return shapeValue["scissors"]
		} else {
			return shapeValue["rock"]
		}
	case outcome == 0:
		if its == "rock" {
			return shapeValue["scissors"]
		} else if its == "paper" {
			return shapeValue["rock"]
		} else {
			return shapeValue["paper"]
		}
	default:
		return shapeValue[its]
	}
}

func getYourOutcome(decision string) int {
	switch decision {
	case "X":
		return 0
	case "Y":
		return 3
	default:
		return 6
	}
}

func getRoundsResult(inputFile *os.File) int {
	scanner := bufio.NewScanner(inputFile)

	var result int
	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Split(line, " ")
		its := opponent[chunks[0]]
		your := you[chunks[1]]
		outcome := resolve(its, your)
		value := shapeValue[your]
		result = result + outcome + value
	}

	return result
}

func resolve(its string, your string) int {
	switch {
	case its == your:
		return 3
	case its == "rock" && your == "paper":
		return 6
	case its == "paper" && your == "scissors":
		return 6
	case its == "scissors" && your == "rock":
		return 6
	default:
		return 0
	}
}
