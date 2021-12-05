package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type world struct {
	m map[coord]int
	i int
}

func newWorld() *world {
	return &world{
		m: make(map[coord]int),
	}
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := newWorld()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		src := parseCoord(parts[0])
		dst := parseCoord(parts[1])

		fillWorld(w, src, dst)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", w.i)
}

func parseCoord(str string) coord {
	parts := strings.Split(str, ",")
	var err error
	var c coord

	c.x, err = strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	c.y, err = strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return c
}

func fillWorld(w *world, src, dst coord) {
	distX, dirX := distAndDir(src.x, dst.x)
	distY, dirY := distAndDir(src.y, dst.y)

	var dist int
	if distX > distY {
		dist = distX
	} else {
		dist = distY
	}

	for i := 0; i <= dist; i++ {
		c := coord{
			src.x + (dirX * i),
			src.y + (dirY * i),
		}

		markWorld(w, c)
	}
}

func markWorld(w *world, c coord) {
	if _, ok := w.m[c]; !ok {
		w.m[c] = 0
	}

	w.m[c]++

	if w.m[c] == 2 {
		w.i++
	}
}

func distAndDir(a, b int) (int, int) {
	var dist, dir int

	if a > b {
		dist = a - b
		dir = -1
	} else if a < b {
		dist = b - a
		dir = 1
	}

	return dist, dir
}
