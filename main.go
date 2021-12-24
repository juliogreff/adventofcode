package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	Unused = 0
	Const  = 1
	Var    = 2
)

func main() {
	path := "2021/24/input_original"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var blocks [][]string
	var block []string
	usage := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			usage = make(map[string]int)
			blocks = append(blocks, block)
			block = nil
		} else {
			if matched, err := regexp.MatchString(`^(div|mul) \w 1$`, line); matched {
				fmt.Printf("line %q optimized away\n", line)
				continue
			} else if err != nil {
				panic(err)
			}

			parts := strings.Split(line, " ")
			op := parts[0]
			lhs := parts[1]
			var rhs string
			if len(parts) > 2 {
				rhs = parts[2]
			}

			if op == "mul" && rhs == "0" && usage[lhs] == Unused {
				fmt.Printf("line %q optimized away\n", line)
				continue
				// omit set to zero
			}

			if op == "add" && isConst(rhs) {
				// if usage[lhs] == Const {
				// 	fmt.Printf("line %q COULD BE optimized away\n", line)
				// }

				usage[lhs] = Const
			} else {
				usage[lhs] = Var
			}

			block = append(block, line)
		}
	}

	for _, b := range blocks {
		printBlock(b)
	}

}

func isConst(op string) bool {
	return op != "x" && op != "y" && op != "z" && op != "w"
}

func printBlock(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
	fmt.Println()
}
