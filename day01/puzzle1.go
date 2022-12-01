package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  inputFile, err := os.Open("./input")
  if err != nil {
    fmt.Println(err)
    return
  }

  bags, err := getBags(inputFile)
  if err != nil {
    fmt.Println(err)
    return
  }

  // fmt.Println(getGreatestBagCalories(bags))
  fmt.Println(getTheThreeGreatestBagCalories(bags))
}

func getTheThreeGreatestBagCalories(bags [][]int) int {
  var first int
  var second int
  var third int
  for _, bag := range bags {
    var sum int
    for _, calories := range bag {
      sum += calories
    }
    if sum > first {
      third = second
      second = first
      first = sum
    } else if sum > second {
      third = second
      second = sum
    } else if sum > third {
      third = sum
    }
  }

  return first + second + third
}

func getGreatestBagCalories(bags [][]int) int {
  var greatest int
  for _, bag := range bags {
    var sum int
    for _, calories := range bag {
      sum += calories
    }
    if sum > greatest {
      greatest = sum
    }
  }

  return greatest
}

func getBags(inputFile *os.File) ([][]int, error) {
  scann := bufio.NewScanner(inputFile)

  bags := [][]int{}
  bag := []int{}
  for scann.Scan() {
    line := strings.Trim(scann.Text(), " ")

    if len(line) > 0 {
      calories, err := strconv.ParseInt(line, 10, 32)
      if err != nil {
        return [][]int{}, err
      }
      
      bag = append(bag, int(calories))
    } else {
      bags = append(bags, bag)
      bag = []int{}
    }
  }

  return bags, nil
}
