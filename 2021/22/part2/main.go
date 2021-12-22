package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type pos struct {
	x int
	y int
	z int
}

type cube struct {
	c1 pos
	c2 pos
}

func (c cube) volume() int {
	return (c.c2.x - c.c1.x + 1) *
		(c.c2.y - c.c1.y + 1) *
		(c.c2.z - c.c1.z + 1)
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cubes := make(map[cube]bool)
	re := regexp.MustCompile(`(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())

		onOff := parseBool(matches[1])
		minX := parseInt(matches[2])
		maxX := parseInt(matches[3])
		minY := parseInt(matches[4])
		maxY := parseInt(matches[5])
		minZ := parseInt(matches[6])
		maxZ := parseInt(matches[7])

		cube1 := cube{
			c1: pos{minX, minY, minZ},
			c2: pos{maxX, maxY, maxZ},
		}

		for cube2 := range cubes {
			if cube1.c1.x > cube2.c2.x || cube1.c2.x < cube2.c1.x || cube1.c1.y > cube2.c2.y || cube1.c2.y < cube2.c1.y || cube1.c1.z > cube2.c2.z || cube1.c2.z < cube2.c1.z {
				continue
			}

			delete(cubes, cube2)

			if cube1.c1.x > cube2.c1.x {
				cubes[cube{
					c1: pos{cube2.c1.x, cube2.c1.y, cube2.c1.z},
					c2: pos{cube1.c1.x - 1, cube2.c2.y, cube2.c2.z},
				}] = true
			}

			if cube1.c2.x < cube2.c2.x {
				cubes[cube{
					c1: pos{cube1.c2.x + 1, cube2.c1.y, cube2.c1.z},
					c2: pos{cube2.c2.x, cube2.c2.y, cube2.c2.z},
				}] = true
			}

			if cube1.c1.y > cube2.c1.y {
				cubes[cube{
					c1: pos{max(cube1.c1.x, cube2.c1.x), cube2.c1.y, cube2.c1.z},
					c2: pos{min(cube1.c2.x, cube2.c2.x), cube1.c1.y - 1, cube2.c2.z},
				}] = true
			}

			if cube1.c2.y < cube2.c2.y {
				cubes[cube{
					c1: pos{max(cube1.c1.x, cube2.c1.x), cube1.c2.y + 1, cube2.c1.z},
					c2: pos{min(cube1.c2.x, cube2.c2.x), cube2.c2.y, cube2.c2.z},
				}] = true

			}

			if cube1.c1.z > cube2.c1.z {
				cubes[cube{
					c1: pos{max(cube1.c1.x, cube2.c1.x), max(cube1.c1.y, cube2.c1.y), cube2.c1.z},
					c2: pos{min(cube1.c2.x, cube2.c2.x), min(cube1.c2.y, cube2.c2.y), cube1.c1.z - 1},
				}] = true
			}

			if cube1.c2.z < cube2.c2.z {
				cubes[cube{
					c1: pos{max(cube1.c1.x, cube2.c1.x), max(cube1.c1.y, cube2.c1.y), cube1.c2.z + 1},
					c2: pos{min(cube1.c2.x, cube2.c2.x), min(cube1.c2.y, cube2.c2.y), cube2.c2.z},
				}] = true
			}
		}

		if onOff {
			cubes[cube1] = true
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var answer int
	for c := range cubes {
		answer += c.volume()
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
