package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

type node struct {
	v    string
	next *node
}

func slow(scanner *bufio.Scanner) {
	var polymer *node
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var last *node
		for _, c := range strings.Split(line, "") {
			if polymer == nil {
				polymer = &node{c, nil}
				last = polymer
			} else {
				last.next = &node{c, nil}
				last = last.next
			}
		}
	}

	rules := make(map[string]map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		pair := strings.Split(parts[0], "")
		insertion := parts[1]

		if _, ok := rules[pair[0]]; !ok {
			rules[pair[0]] = make(map[string]string)
		}

		rules[pair[0]][pair[1]] = insertion
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for step := 0; step < Steps; step++ {
		p := polymer
		for p.next != nil {
			_, ok := rules[p.v]
			if !ok {
				p = p.next
				continue
			}

			insert, ok := rules[p.v][p.next.v]
			if !ok {
				p = p.next
				continue
			}

			nxt := p.next
			p.next = &node{v: insert, next: nxt}
			p = nxt
		}
	}

	min := math.MaxInt
	max := 0

	counters := make(map[string]int)
	p := polymer
	for p != nil {
		counters[p.v]++
		p = p.next
	}

	for p, c := range counters {
		_ = p
		if c < min {
			min = c
		} else if c > max {
			max = c
		}
	}

	fmt.Printf("%d\n", max-min)
}

func print(p *node) {
	for p != nil {
		fmt.Print(p.v)
		p = p.next
	}
	fmt.Println()
}
