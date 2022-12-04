package file

import (
	"bufio"
	"io"
	"io/fs"
)

type input struct {
  lines []string
}

func (i input) GetContents(fileSystem fs.FS, filename string) ([]string, error) {
  inputFile, err := fileSystem.Open(filename)
  if err != nil {
    return []string{}, err
  }
  defer inputFile.Close()

  i.getLines(inputFile)

  return i.lines, nil
}

func (i *input) getLines(inputFile io.Reader) {
  scanner := bufio.NewScanner(inputFile)

  for scanner.Scan() {
    i.lines = append(i.lines, scanner.Text())
  }
}

func NewReader() *input {
  return &input{}
}
