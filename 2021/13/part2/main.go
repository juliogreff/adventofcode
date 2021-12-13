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
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var maxX, maxY int
	for dot := range paper {
		if dot.x > maxX {
			maxX = dot.x
		}

		if dot.y > maxY {
			maxY = dot.y
		}
	}

	draw := make([][]bool, maxY+1)
	for dot := range paper {
		if draw[dot.y] == nil {
			draw[dot.y] = make([]bool, maxX+1)
		}

		draw[dot.y][dot.x] = true
	}

	for y := range draw {
		for x := range draw[y] {
			if draw[y][x] {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
