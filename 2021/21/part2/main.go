package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const MaxScore = 21

type state struct {
	p1 int
	p2 int
	s1 int
	s2 int
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`Player \d starting position: (\d)`)
	var players []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		pos := parseInt(matches[1])
		players = append(players, pos)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	possibleRolls := make(map[int]int)
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				possibleRolls[a+b+c]++
			}
		}
	}

	st := state{players[0], players[1], 0, 0}
	universes := map[state]float64{
		st: 1,
	}

	turn := 0
	more := true

	for more {
		more = false
		next := make(map[state]float64)

		for u, uCount := range universes {
			if u.s1 < MaxScore && u.s2 < MaxScore {
				more = true
				for roll, rCount := range possibleRolls {
					s := u
					if (turn % 2) == 0 {
						s.p1 = (s.p1-1+roll)%10 + 1
						s.s1 += s.p1
					} else {
						s.p2 = (s.p2-1+roll)%10 + 1
						s.s2 += s.p2
					}

					next[s] += uCount * float64(rCount)
				}
			} else {
				next[u] += uCount
			}
		}

		universes = next
		turn++
	}

	var p1, p2 float64
	for u, c := range universes {
		if u.s1 >= MaxScore {
			p1 += c
		} else if u.s2 >= MaxScore {
			p2 += c
		} else {
			panic("there's an universe where a player hasn't won")
		}
	}

	fmt.Printf("%f\n", math.Max(p1, p2))
}

func copyUniverses(old map[state]int) map[state]int {
	cp := make(map[state]int)
	for k, v := range old {
		cp[k] = v
	}
	return cp
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
