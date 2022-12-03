package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var elf, answer int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if elf > answer {
				answer = elf
			}
			elf = 0
			continue
		}

		calories, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}

		elf += calories
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}
