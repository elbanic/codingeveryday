package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	newNode := helper(node, make(map[int]*Node))
	return newNode
}

func helper(node *Node, m map[int]*Node) *Node {
	if node == nil {
		return node
	}

	if v, ok := m[node.Val]; ok {
		return v
	}

	newNode := &Node{Val: node.Val}
	m[node.Val] = newNode

	var neighbors []*Node
	for _, v := range node.Neighbors {
		neighbors = append(neighbors, helper(v, m))
	}
	newNode.Neighbors = neighbors
	return newNode
}

func createGraph(list [][]int) *Node {

	if len(list) == 0 {
		return nil
	}
	first := &Node{}
	cur := first
	m := make(map[int]*Node)
	m[1] = first

	for i := range list {
		cur.Val = i + 1
		var nodes []*Node
		for j := range list[i] {
			if node, ok := m[list[i][j]]; ok {
				nodes = append(nodes, node)
				cur.Neighbors = nodes
			} else {
				node = &Node{Val: list[i][j]}
				nodes = append(nodes, node)
				cur.Neighbors = nodes
				m[list[i][j]] = node
			}
		}
		cur = m[i+2]
	}
	return first
}

func (node *Node) print() {

	queue := []*Node{node}
	visited := make(map[int]bool)
	visited[node.Val] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Print(cur.Val, "->")
		for i, n := range cur.Neighbors {
			fmt.Print(n.Val, n)
			if i != len(cur.Neighbors)-1 {
				fmt.Print(",")
			}
		}
		fmt.Println()
		for _, n := range cur.Neighbors {
			if !visited[n.Val] {
				queue = append(queue, n)
				visited[n.Val] = true
			}
		}
	}
}

func main() {
	adjList := [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}
	graph := createGraph(adjList)
	graph.print()
	fmt.Println()

	CloneGraph := cloneGraph(graph)
	CloneGraph.print()

	cloneGraph2 := cloneGraph(createGraph([][]int{}))
	cloneGraph2.print()
}
