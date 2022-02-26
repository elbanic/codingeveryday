package main

import (
	"fmt"
	"math"
)

const MaxInt = int(^uint32(0) >> 1)

var cache [][]int

func shortestPathLength(graph [][]int) int {

	n := len(graph)
	endingMask := (1 << n) - 1
	cache = make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, endingMask+1)
	}
	best := MaxInt
	for node := 0; node < n; node++ {
		best = int(math.Min(float64(best), float64(dp(node, endingMask, graph))))
	}
	return best
}

func dp(node, mask int, graph [][]int) int {

	if cache[node][mask] != 0 {
		return cache[node][mask]
	}
	if (mask & (mask - 1)) == 0 {
		return 0
	}

	cache[node][mask] = MaxInt - 1
	for _, neighbor := range graph[node] {
		if (mask & (1 << neighbor)) != 0 {
			visited := dp(neighbor, mask, graph)
			notVisited := dp(neighbor, mask^(1<<node), graph)
			betterOption := math.Min(float64(visited), float64(notVisited))
			cache[node][mask] = int(math.Min(float64(cache[node][mask]), 1+betterOption))
		}
	}
	return cache[node][mask]
}

func main() {
	graph := [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}
	fmt.Println(shortestPathLength(graph))
}
