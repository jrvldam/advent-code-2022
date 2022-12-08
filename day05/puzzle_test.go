package main_test

import (
	"bufio"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestGetLines(t *testing.T) {
  content := getContent()

  scanner := bufio.NewScanner(content)
  scanner.Scan()
  got := scanner.Text()
  want := "    [D]    "

  if got != want {
    t.Errorf("want %q, got %q", want, got)
  }
}

func getContent() fs.File {
  file := fstest.MapFS{
    "input": {Data: []byte(inputContents)},
  }

  content, err := file.Open("input")
  if err != nil {
    panic(err)
  }
  
  return content
}

var inputContents = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
