package graphs

import (
	"math"

	"github.com/juliogreff/adventofcode/pkg/xy"
)

type DijkstraOpts struct {
	Cost func(graph [][]int, src xy.XY, dst xy.XY) int
}

func DefaultDijkstraOpts() DijkstraOpts {
	return DijkstraOpts{
		Cost: func(graph [][]int, src xy.XY, dst xy.XY) int {
			return graph[dst.Y][dst.X]
		},
	}
}

func Dijkstra(graph [][]int, src xy.XY, dst xy.XY, opts DijkstraOpts) int {
	distances := make(map[xy.XY]int)
	distances[src] = 0

	visited := make(map[xy.XY]struct{})
	unvisited := make(map[xy.XY]struct{})
	for _, n := range neighbors(graph, src) {
		unvisited[n] = struct{}{}
	}

	for len(unvisited) > 0 {
		delete(unvisited, src)

		for _, n := range neighbors(graph, src) {
			if _, ok := visited[n]; ok {
				continue
			}

			cost := opts.Cost(graph, src, n)
			if cost == math.MaxInt {
				continue
			}

			dtn := distances[src] + cost

			d, ok := distances[n]
			if !ok {
				d = math.MaxInt
			}

			if dtn < d {
				distances[n] = dtn
			}
		}

		visited[src] = struct{}{}

		if src == dst {
			break
		}

		src = getClosestUnvisited(graph, src, visited, unvisited, distances)
	}

	return distances[dst]
}

func getClosestUnvisited(
	graph [][]int,
	src xy.XY,
	visited map[xy.XY]struct{},
	unvisited map[xy.XY]struct{},
	distances map[xy.XY]int,
) xy.XY {
	var node xy.XY
	lowest := math.MaxInt

	neighbors := neighbors(graph, src)
	for _, n := range neighbors {
		if _, ok := visited[n]; ok {
			continue
		}

		unvisited[n] = struct{}{}
	}

	for n := range unvisited {
		if d, ok := distances[n]; ok && d < lowest {
			node = n
			lowest = distances[n]
		}
	}

	return node
}

func neighbors(graph [][]int, src xy.XY) []xy.XY {
	var n []xy.XY

	if src.X+1 < len(graph[0]) {
		n = append(n, xy.XY{src.X + 1, src.Y})
	}

	if src.X >= 1 {
		n = append(n, xy.XY{src.X - 1, src.Y})
	}

	if src.Y+1 < len(graph) {
		n = append(n, xy.XY{src.X, src.Y + 1})
	}

	if src.Y >= 1 {
		n = append(n, xy.XY{src.X, src.Y - 1})
	}

	return n
}
