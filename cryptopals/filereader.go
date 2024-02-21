package cryptopals

import (
	"bufio"
	"os"
)

// Works for fastening file reading structures
func ReadFile(path *string) (fileLines []string) {

	// Check is there's a valid path
	file, ferr := os.Open(*path)
	if ferr != nil {
		panic(ferr)
	}

	// If path is valid, we insert all lines into a slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileLines = append(fileLines, line)
	}

	// As we declared our return in the func structure
	// there is no need to rewrite it here
	return
}
