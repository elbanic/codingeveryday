/*
https://leetcode.com/problems/all-paths-from-source-to-target/
Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1,
find all possible paths from node 0 to node n - 1 and return them in any order.
The graph is given as follows: graph[i] is a list of all nodes you can visit from node i
(i.e., there is a directed edge from node i to node graph[i][j]).

Example 1:
	Input: graph = [[1,2],[3],[3],[]]
	Output: [[0,1,3],[0,2,3]]
	Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.

Example 2:
	Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
	Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

Example 3:
	Input: graph = [[1],[]]
	Output: [[0,1]]

Example 4:
	Input: graph = [[1,2,3],[2],[3],[]]
	Output: [[0,1,2,3],[0,2,3],[0,3]]

Example 5:
	Input: graph = [[1,3],[2],[3],[]]
	Output: [[0,1,2,3],[0,3]]
*/
package main

import "fmt"

//dfs
func allPathsSourceTarget(graph [][]int) [][]int {

	_, output := findNode(graph, []int{0}, []int{}, [][]int{})

	return output
}

func findNode(graph [][]int, stack []int, path []int, output [][]int) ([]int, [][]int){
	if len(stack) == 0 {
		return stack, output
	}
	cur := stack[len(stack) - 1]
	path = append(path, cur)

	if cur == len(graph) - 1 {
		output = append(output, path)
		return stack, output
	}

	stack = stack[:len(stack) - 1]
	for _,v := range graph[cur] {
		stack = append(stack, v)
		temp := make([]int, len(path))
		copy(temp, path)
		stack, output = findNode(graph, stack, temp, output)
	}
	return stack, output
}

func main() {

	graph := [][]int{{3,1},{4,6,7,2,5},{4,6,3},{6,4},{7,6,5},{6},{7},{}}
	fmt.Println(allPathsSourceTarget(graph))
}
