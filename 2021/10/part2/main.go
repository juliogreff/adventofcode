package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

var illegalPoints = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incompletePoints = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var scores []int

line:
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
					continue line
				}
			}
		}

		var score int
		for i := len(stack) - 1; i >= 0; i-- {
			c := mapping[stack[i]]
			score = score*5 + incompletePoints[c]
		}

		scores = append(scores, score)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(scores)

	fmt.Printf("%d\n", scores[len(scores)/2])
}
