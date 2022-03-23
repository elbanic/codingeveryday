package main

import "fmt"

func minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		if _, ok := graph[arr[i]]; ok {
			graph[arr[i]] = append(graph[arr[i]], i)
		} else {
			graph[arr[i]] = []int{}
			graph[arr[i]] = append(graph[arr[i]], i)
		}
	}

	var cur []int
	cur = append(cur, 0)
	visited := make(map[int]bool)

	step := 0
	for len(cur) > 0 {
		var next []int

		for _, node := range cur {
			if node == n-1 {
				return step
			}
			for _, child := range graph[arr[node]] {
				if !visited[child] {
					visited[child] = true
					next = append(next, child)
				}
			}
			graph[arr[node]] = []int{}

			if node+1 < n && !visited[node+1] {
				visited[node+1] = true
				next = append(next, node+1)
			}
			if node-1 >= 0 && !visited[node-1] {
				visited[node-1] = true
				next = append(next, node-1)
			}
		}
		cur = next
		step++
	}
	return -1
}

func main() {
	arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	fmt.Println(minJumps(arr))
}
