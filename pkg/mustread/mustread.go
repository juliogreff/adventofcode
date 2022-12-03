package mustread

import (
	"bufio"
	"os"
)

func File(path string, fn func(*bufio.Scanner)) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fn(scanner)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
