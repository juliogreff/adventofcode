package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type boardLine []int
type board []boardLine

const boardSize = 5

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ok := scanner.Scan()
	if !ok {
		panic("nothing to read!")
	}

	strs := strings.Split(scanner.Text(), ",")
	numbers := make([]int, 0, len(strs))
	for _, str := range strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, n)
	}

	var boards []board
	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			boards = append(boards, board{})
			continue
		}

		b := len(boards) - 1
		boards[b] = append(boards[b], parseBoardLine(txt))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var sum, n int
	picked := make(map[int]struct{})
	for _, n = range numbers {
		picked[n] = struct{}{}

		sum = checkWinners(boards, picked)
		if sum > 0 {
			break
		}
	}

	fmt.Printf("%d %d %d\n", sum, n, sum*n)
}

func checkWinners(boards []board, picked map[int]struct{}) int {
	for _, b := range boards {
		sum := boardComplete(b, picked)
		if sum > 0 {
			fmt.Printf("%+v\n", b)
			return sum
		}
	}

	return 0
}

func boardComplete(b board, picked map[int]struct{}) int {
	var (
		sum      int
		complete bool
	)
	for d := 0; d < boardSize; d++ {
		lineOk := true

		for i := 0; i < boardSize; i++ {
			n := b[d][i]
			_, ok := picked[n]
			if ok {
				// do nothing
			} else {
				lineOk = false
				sum += n
			}
		}

		if complete || lineOk {
			complete = true
			continue
		}

		columnOk := true
		for j := 0; j < boardSize; j++ {
			n := b[j][d]
			_, ok := picked[n]
			if !ok {
				columnOk = false
				break
			}
		}

		if columnOk {
			complete = true
		}
	}

	if complete {
		return sum
	} else {
		return 0
	}
}

func parseBoardLine(str string) boardLine {
	strs := strings.Split(str, " ")
	line := make(boardLine, 0, boardSize)

	for _, s := range strs {
		if len(s) == 0 {
			continue
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		line = append(line, n)
	}

	return line
}
