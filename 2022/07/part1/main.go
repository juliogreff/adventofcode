package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
)

const MinSize = 100000

type dir struct {
	size int
	tree tree
}

type tree map[string]*dir

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		path := []string{}
		t := make(tree)

		for scanner.Scan() {
			line := scanner.Text()

			parts := strings.Split(line, " ")

			switch parts[0] {
			case "$":
				switch parts[1] {
				case "cd":
					if parts[2] == ".." {
						path = path[:len(path)-1]
					} else {
						path = append(path, parts[2])
					}

				case "ls":
					continue
				}

			case "dir":

			default:
				addSize(t, path, mustparse.Int(parts[0]))
			}
		}

		fmt.Printf("%d\n", walkTree(t))

	})
}

func walkTree(t tree) int {
	var total int
	for _, d := range t {
		if d.size < MinSize {
			total += d.size
		}

		if d.tree != nil {
			total += walkTree(d.tree)
		}
	}

	return total
}

func addSize(t tree, path []string, size int) {
	for _, p := range path {
		if t[p] == nil {
			t[p] = &dir{
				tree: make(tree),
			}
		}

		t[p].size += size
		t = t[p].tree
	}
}
