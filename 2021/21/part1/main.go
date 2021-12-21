package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`Player \d starting position: (\d)`)
	var players []int
	var scores []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		pos := parseInt(matches[1])
		players = append(players, pos)
		scores = append(scores, 0)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var i int
	var roll int
	for {
		var d int
		for j := 0; j < 3; j++ {
			roll++
			if roll > 100 {
				roll = 1
			}
			d += roll
		}

		p := i % len(players)
		players[p] += d
		players[p] = players[p] % 10
		if players[p] == 0 {
			players[p] = 10
		}

		scores[p] += players[p]

		i++

		if scores[p] >= 1000 {
			break
		}
	}

	fmt.Printf("%d\n", i*3*scores[i%2])
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
