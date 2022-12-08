package file

import (
	"bufio"
	"os"
)

func GetLinesFromFile(file *os.File) (lines []string) {
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines
}

func GetFile(path string) *os.File {
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  return file
}
