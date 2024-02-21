package cryptopals

import (
	"bufio"
	"os"
)

func ReadFile(path string) []string {
	file, ferr := os.Open(path)
	if ferr != nil {
		panic(ferr)
	}

	var fileLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileLines = append(fileLines, line)
	}
	
	return fileLines
}
