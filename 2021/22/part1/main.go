package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type pos struct {
	x int
	y int
	z int
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reactor := make(map[pos]bool)
	re := regexp.MustCompile(`(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)

	var answer int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())

		onOff := parseBool(matches[1])
		minX := clampMin(parseInt(matches[2]))
		maxX := clampMax(parseInt(matches[3]))
		minY := clampMin(parseInt(matches[4]))
		maxY := clampMax(parseInt(matches[5]))
		minZ := clampMin(parseInt(matches[6]))
		maxZ := clampMax(parseInt(matches[7]))

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				for z := minZ; z <= maxZ; z++ {
					p := pos{x, y, z}
					prev := reactor[p]

					if prev && !onOff {
						answer--
					} else if !prev && onOff {
						answer++
					}

					reactor[p] = onOff
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}

func parseInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func parseBool(str string) bool {
	switch str {
	case "on":
		return true
	case "off":
		return false
	}

	panic("not a bool")
}

func clampMin(i int) int {
	if i < -50 {
		return -50
	}

	return i
}

func clampMax(i int) int {
	if i > 50 {
		return 50
	}

	return i
}
