package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var x, y int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		dir := parts[0]
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch dir {
		case "forward":
			x += amount
		case "up":
			y -= amount
		case "down":
			y += amount
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", x*y)
}
