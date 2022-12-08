package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jrvldam/advent-code-2022/shared/stack"
)

type Procedure struct {
	quantity int
	from     int
	to       int
}

func main() {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	lines := GetLines(inputFile)
	start, steps := processInput(lines)

	store := createStore(start)
	procedures := createProcedures(steps)

	// applyProceduresCrateMover9000(&store, procedures)
	applyProceduresCrateMover9001(&store, procedures)

	keySlice := make([]int, 0)
	for idx := range store {
		keySlice = append(keySlice, idx)
	}
	sort.Ints(keySlice)

	for _, idx := range keySlice {
		pile := store[idx]
		crate := pile.Pick()
		fmt.Printf(crate)
	}
	fmt.Println()
}

func applyProceduresCrateMover9001(store *Store, procedures []Procedure) {
	for _, procedure := range procedures {
		source, err := store.Search(procedure.from)
		if err != nil {
			panic(err)
		}

		target, err := store.Search(procedure.to)
		if err != nil {
			panic(err)
		}

		quantity := procedure.quantity

		var upload []string
		for i := 0; i < quantity; i += 1 {
			crate := source.Pop()
			upload = append(upload, crate)
		}

		for i := len(upload) - 1; i >= 0; i -= 1 {
			target.Push(upload[i])
		}
	}
}

func applyProceduresCrateMover9000(store *Store, procedures []Procedure) {
	for _, procedure := range procedures {
		source, err := store.Search(procedure.from)
		if err != nil {
			panic(err)
		}

		target, err := store.Search(procedure.to)
		if err != nil {
			panic(err)
		}

		quantity := procedure.quantity

		for i := 0; i < quantity; i += 1 {
			crate := source.Pop()
			target.Push(crate)
		}
	}
}

func createProcedures(lines []string) []Procedure {
	var procedures []Procedure
	for _, line := range lines {
		chunks := strings.Split(line, " ")
		quantity, _ := strconv.ParseInt(chunks[1], 10, 32)
		from, _ := strconv.ParseInt(chunks[3], 10, 32)
		to, _ := strconv.ParseInt(chunks[5], 10, 32)
		procedure := Procedure{quantity: int(quantity), from: int(from), to: int(to)}
		procedures = append(procedures, procedure)
	}

	return procedures
}

func createStore(lines []string) Store {
	regxLine := regexp.MustCompile(".{3,4}")
	regxToken := regexp.MustCompile("[A-Z]")
	regxNumbers := regexp.MustCompile("\\d")

	store := NewStore()

	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	for _, line := range lines {
		chunks := regxLine.FindAll([]byte(line), -1)
		number := regxNumbers.Find(chunks[0])

		if number != nil {
			continue
		}

		for idx, chunk := range chunks {
			crate := regxToken.Find(chunk)
			if crate != nil {
				store.Add(idx+1, stack.NewStack())
				pile, _ := store.Search(idx + 1)
				pile.Push(string(crate))
			}
		}
	}

	return store
}

func processInput(lines []string) ([]string, []string) {
	flag := false

	var start []string
	var instructions []string
	for _, line := range lines {
		isEmptyLine := strings.TrimSpace(line) == ""

		if !flag && isEmptyLine {
			flag = true
			continue
		}

		if flag {
			instructions = append(instructions, line)
		} else {
			start = append(start, line)
		}
	}

	return start, instructions
}

func GetLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
