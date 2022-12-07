package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
)

const (
	SpaceAvailable = 70000000
	MinimumUnused  = 30000000
)

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

		fmt.Printf("%d\n", walkTree(t["/"], t))
	})
}

func walkTree(root *dir, t tree) int {
	spaceFreed := math.MaxInt64
	for _, d := range t {
		f := root.size - d.size

		if f >= SpaceAvailable-MinimumUnused {
			continue
		}

		if f < spaceFreed {
			spaceFreed = d.size
		}

		if d.tree != nil {
			f = walkTree(root, d.tree)
			if f < spaceFreed {
				spaceFreed = f
			}
		}
	}

	return spaceFreed
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
