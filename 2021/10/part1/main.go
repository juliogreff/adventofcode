package main

import (
	"bufio"
	"fmt"
	"os"
)

var mapping = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var reverse = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var points = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var answer int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var stack []byte
		chunks := []byte(scanner.Text())
		for _, c := range chunks {
			if _, ok := mapping[c]; ok {
				stack = append(stack, c)
			} else {
				last := stack[len(stack)-1]
				if r := reverse[c]; r == last {
					stack = stack[:len(stack)-1]
				} else {
					answer += points[c]
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}
