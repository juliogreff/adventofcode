package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	graph := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := scanner.Text()
		parts := strings.Split(path, "-")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	paths := pathsToEnd(graph, map[string]struct{}{}, []string{"start"})

	answer := len(paths)

	fmt.Printf("%d\n", answer)
}

func pathsToEnd(graph map[string][]string, visited map[string]struct{}, path []string) [][]string {
	var paths [][]string

	src := path[len(path)-1]

	for _, dst := range graph[src] {
		if isLower(dst) && find(path, dst) {
			continue
		}

		p := append([]string{}, path...)
		p = append(p, dst)

		if dst == "end" {
			paths = append(paths, p)
			continue
		}

		paths = append(paths, pathsToEnd(graph, visited, p)...)
	}

	return paths
}

func find(haystack []string, needle string) bool {
	for _, x := range haystack {
		if needle == x {
			return true
		}
	}

	return false
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
