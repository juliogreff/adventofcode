package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	paper := make(map[pos]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		paper[pos{x, y}] = struct{}{}
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		fold := strings.Split(parts[2], "=")
		axis := fold[0]
		where, err := strconv.Atoi(fold[1])
		if err != nil {
			panic(err)
		}

		newPaper := make(map[pos]struct{})

		for dot := range paper {
			var newDot pos

			switch axis {
			case "x":
				if dot.x > where {
					newDot.y = dot.y
					newDot.x = 2*where - dot.x
				} else {
					newDot = dot
				}
			case "y":
				if dot.y > where {
					newDot.x = dot.x
					newDot.y = 2*where - dot.y
				} else {
					newDot = dot
				}
			}

			newPaper[newDot] = struct{}{}
		}

		paper = newPaper

		break
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	answer := len(paper)
	fmt.Printf("%d\n", answer)
}
