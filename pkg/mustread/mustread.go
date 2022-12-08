package mustread

import (
	"bufio"
	"os"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
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

func FileMapOfInts(path string) [][]int {
	var mapOfInts [][]int

	File(os.Args[1], func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			mapOfInts = append(mapOfInts, mustparse.Ints(scanner.Text()))
		}
	})

	return mapOfInts
}
