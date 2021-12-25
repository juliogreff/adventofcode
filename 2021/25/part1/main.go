package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos int
type move struct {
	i int
	j int
	k int
	l int
}

const (
	None pos = iota
	East
	South
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var floor [][]pos

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var p pos
		var l []pos

		for _, c := range line {
			switch c {
			case '.':
				p = None
			case '>':
				p = East
			case 'v':
				p = South
			}

			l = append(l, p)
		}

		floor = append(floor, l)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var step int
	for {
		var moved bool
		for _, which := range []pos{East, South} {
			var moves []move
			for i := range floor {
				for j := range floor[i] {
					if floor[i][j] != which {
						continue
					}

					if m, ok := canMove(floor, which, i, j); ok {
						moved = true
						moves = append(moves, m)
					}
				}
			}

			for _, m := range moves {
				floor[m.i][m.j] = None
				floor[m.k][m.l] = which
			}
		}

		step++

		if !moved {
			break
		}
	}

	fmt.Printf("%d\n", step)
}

func canMove(floor [][]pos, which pos, i int, j int) (move, bool) {
	var k, l int

	switch which {
	case East:
		k = i
		l = j + 1
		if l == len(floor[k]) {
			l = 0
		}
	case South:
		k = i + 1
		l = j

		if k == len(floor) {
			k = 0
		}
	}

	if floor[k][l] == None {
		return move{i, j, k, l}, true
	} else {
		return move{}, false
	}
}

func print(floor [][]pos) {
	for i := range floor {
		for j := range floor[i] {
			switch floor[i][j] {
			case None:
				fmt.Print(".")
			case East:
				fmt.Print(">")
			case South:
				fmt.Print("v")
			}
		}
		fmt.Println()
	}
}
