package file

import (
	"bufio"
	"os"
)

type fileReader struct {
	file *os.File
}

func (f fileReader) GetLinesFromFile(path string) (lines []string) {
	f.getFile(path)
	defer f.file.Close()

	scanner := bufio.NewScanner(f.file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func (f *fileReader) getFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	f.file = file
}

func NewFileReader() fileReader {
	return fileReader{}
}
