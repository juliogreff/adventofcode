package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
)

type monkey struct {
	inspected int
	items     []int
	op        operation
	test      int
	ifTrue    int
	ifFalse   int
}

type operation struct {
	l  *int
	r  *int
	op string
}

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var (
			monkeys []*monkey
		)

		max := 1

		for scanner.Scan() {
			line := strings.Trim(scanner.Text(), " ")

			if strings.HasPrefix(line, "Monkey") {
				monkeys = append(monkeys, &monkey{})
				continue
			}

			monkey := monkeys[len(monkeys)-1]
			parts := strings.Split(line, ": ")

			switch parts[0] {
			case "Starting items":
				monkey.items = mustparse.SplitInts(parts[1], ", ")
			case "Operation":
				var l, r, op string

				fmt.Sscanf(parts[1], "new = %s%s%s", &l, &op, &r)

				monkey.op = operation{
					op: op,
				}

				if l != "old" {
					l := mustparse.Int(l)
					monkey.op.l = &l
				}

				if r != "old" {
					r := mustparse.Int(r)
					monkey.op.r = &r
				}
			case "Test":
				fmt.Sscanf(parts[1], "divisible by %d", &monkey.test)
				max *= monkey.test
			case "If true":
				fmt.Sscanf(parts[1], "throw to monkey %d", &monkey.ifTrue)
			case "If false":
				fmt.Sscanf(parts[1], "throw to monkey %d", &monkey.ifFalse)
			case "":
				continue
			default:
				panic(fmt.Sprintf("Operation %q not known", parts[0]))
			}
		}

		rounds := 10000
		for round := 0; round < rounds; round++ {
			for _, m := range monkeys {
				for _, item := range m.items {
					m.inspected++
					n := doOperation(item, m.op)

					if n%m.test == 0 {
						monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, n%max)
					} else {
						monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, n%max)
					}
				}

				m.items = nil
			}
		}

		var allInspected []int
		for _, m := range monkeys {
			allInspected = append(allInspected, m.inspected)
		}

		sort.Ints(allInspected)

		fmt.Printf("%d\n", allInspected[len(allInspected)-1]*allInspected[len(allInspected)-2])
	})
}

func doOperation(i int, op operation) int {
	var l, r int

	if op.l == nil {
		l = i
	} else {
		l = *op.l
	}

	if op.r == nil {
		r = i
	} else {
		r = *op.r
	}

	switch op.op {
	case "+":
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	case "/":
		return l / r
	default:
		panic("not reachable")
	}
}
