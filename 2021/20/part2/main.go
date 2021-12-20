package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Passes = 50
const Bleed = 2

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var algo []bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		algo = parseLine(line)
	}

	var img [][]bool
	for scanner.Scan() {
		line := scanner.Text()
		img = append(img, parseLine(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := 0; i < Passes; i++ {
		background := false

		if algo[0] {
			background = i%2 != 0
		}

		img = enhance(img, algo, background)
	}

	print(img)

	answer := countLitPixels(img)

	fmt.Printf("%d\n", answer)
}

func enhance(img [][]bool, algo []bool, background bool) [][]bool {
	var newImg [][]bool

	for i := -Bleed; i < len(img)+Bleed; i++ {
		var line []bool
		for j := -Bleed; j < len(img[0])+Bleed; j++ {
			idx := algoIndex(img, i, j, background)
			line = append(line, algo[idx])
		}
		newImg = append(newImg, line)
	}

	return newImg
}

func algoIndex(img [][]bool, i, j int, bg bool) int {
	var str string

	background := "0"
	if bg {
		background = "1"
	}

	for a := i - 1; a <= i+1; a++ {
		for b := j - 1; b <= j+1; b++ {
			if a < 0 || b < 0 {
				str += background
				continue
			}

			if a >= len(img) {
				str += background
				continue
			}

			if b >= len(img[a]) {
				str += background
				continue
			}

			if img[a][b] {
				str += "1"
			} else {
				str += "0"
			}
		}
	}

	idx, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(idx)
}

func countLitPixels(img [][]bool) int {
	var c int
	for i := range img {
		for _, pixel := range img[i] {
			if pixel {
				c++
			}
		}
	}
	return c
}

func print(img [][]bool) {
	fmt.Println()
	for i := range img {
		for _, pixel := range img[i] {
			if pixel {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseLine(str string) []bool {
	var parsed []bool
	for _, c := range str {
		parsed = append(parsed, c == '#')
	}
	return parsed
}
