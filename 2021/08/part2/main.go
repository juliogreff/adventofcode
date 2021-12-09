package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

const (
	One   = "cf"
	Seven = "acf"
	Four  = "bcdf"
	Two   = "acdeg"
	Three = "acdfg"
	Five  = "abdfg"
	Six   = "abdefg"
	Nine  = "abcdfg"
	Zero  = "abcefg"
	Eight = "abcdefg"
)

var numbers = []string{
	Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine,
}

var numbersByLength = map[int][]string{
	2: []string{One},
	3: []string{Seven},
	4: []string{Four},
	5: []string{Two, Three, Five},
	6: []string{Zero, Six, Nine},
	7: []string{Eight},
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var answer int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")

		patterns := strings.Split(parts[0], " ")
		digits := strings.Split(parts[1], " ")

		mapping := findMapping(patterns)

		answer += translateOutput(mapping, digits)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}

func findMapping(patterns []string) map[byte][]byte {
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i]) < len(patterns[j])
	})

	mapping := make(map[byte][]byte)
	takenTargets := make(map[byte]struct{})

pattern:
	for _, patternStr := range patterns {
		pattern := []byte(patternStr)

		var notTaken []byte
		taken := make(map[byte]struct{})
		for k, v := range takenTargets {
			taken[k] = v
		}

		m := make(map[byte][]byte)
		for k, v := range mapping {
			m[k] = v
		}

		for _, numberStr := range numbersByLength[len(pattern)] {
			number := []byte(numberStr)
			for _, tgt := range number {
				if _, ok := taken[tgt]; !ok {
					taken[tgt] = struct{}{}
					notTaken = append(notTaken, tgt)
				}
			}

			if len(pattern) <= 4 {
				for _, src := range pattern {
					if len(m[src]) == 0 {
						m[src] = notTaken
					}
				}
			} else if len(pattern) == 5 {
				for _, src := range []byte(Eight) {
					t := m[src]
					if len(t) == 0 {
						m[src] = notTaken
					}
				}
			}

			if isPossible(m, number, pattern) {
				for _, c := range pattern {
					m[c] = intersection(m[c], number)
				}

				mapping = m
				takenTargets = taken

				continue pattern
			}
		}
	}

	dedupe(mapping)

	return mapping
}

func translateOutput(mapping map[byte][]byte, digits []string) int {
	var n int
	for i, d := range digits {
		d = translateDigit(mapping, d)

		split := strings.Split(d, "")
		sort.Strings(split)
		d = strings.Join(split, "")

		var (
			m   int
			str string
		)
		for m, str = range numbers {
			if str == d {
				break
			}
		}

		n += int(math.Pow10(len(digits)-1-i)) * m
	}

	return n
}

func translateDigit(mapping map[byte][]byte, n string) string {
	str := ""
	for _, x := range []byte(n) {
		w := mapping[x]
		str += string(w[0])

		if len(w) > 1 {
			print(mapping)
			panic("yikes")
		}
	}

	return str
}

func print(mapping map[byte][]byte) {
	for _, i := range []byte(Eight) {
		fmt.Printf("%s => [%s] ", string(i), mapping[i])
	}
	fmt.Println()
}

func isPossible(mapping map[byte][]byte, n []byte, pattern []byte) bool {
	seen := make(map[byte]struct{})

	for _, c := range n {
		var found bool
		for _, m := range []byte(Eight) {
			chars, ok := mapping[m]
			if !ok || !find(pattern, m) {
				continue
			}

			if find(chars, c) {
				if _, ok := seen[m]; ok {
					continue
				}

				seen[m] = struct{}{}
				found = true

				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func difference(a, b []byte) []byte {
	var ret []byte
	for _, x := range a {
		if !find(b, x) {
			ret = append(ret, x)
		}
	}

	return ret
}

func intersection(a, b []byte) []byte {
	var ret []byte
	for _, x := range a {
		if find(b, x) {
			ret = append(ret, x)
		}
	}

	return ret
}

func find(s []byte, b byte) bool {
	for _, x := range s {
		if x == b {
			return true
		}
	}

	return false
}

func dedupe(mapping map[byte][]byte) {
outer:
	for s, targets := range mapping {
		if len(targets) > 1 {
			for i, t := range targets {
				for _, targets := range mapping {
					if len(targets) == 1 && find(targets, t) {
						l := len(mapping[s])
						mapping[s][i] = mapping[s][l-1]
						mapping[s] = mapping[s][:l-1]
						continue outer
					}
				}
			}
		}
	}
}
