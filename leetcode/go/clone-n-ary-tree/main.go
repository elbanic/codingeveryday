package main

type Node struct {
	Val      int
	Children []*Node
}

func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	ret := traversalNaryTree(root)
	return createNaryTree(ret)
}

func traversalNaryTree(root *Node) []int {

	queue := []*Node{root}
	ret := []int{root.Val, -1}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, child := range cur.Children {
			queue = append(queue, child)
			ret = append(ret, child.Val)
		}

		ret = append(ret, -1)
	}
	return ret
}

func createNaryTree(nums []int) *Node {

	if len(nums) == 0 {
		return nil
	}

	root := &Node{Val: nums[0]}
	queue := []*Node{root}
	var children []*Node
	for i := 2; i < len(nums); i++ {
		if nums[i] != -1 {
			cur := &Node{Val: nums[i]}
			children = append(children, cur)
			queue = append(queue, cur)
		} else {
			parent := queue[0]
			queue = queue[1:]
			parent.Children = children
			children = nil
		}
	}
	if len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		parent.Children = children
	}

	return root
}

func main() {
	root := []int{1, -1, 3, 2, 4, -1, 5, 6}
	ntree := createNaryTree(root)
	cloneTree(ntree)

	root2 := []int{1, -1, 2, 3, 4, 5, -1, -1, 6, 7, -1, 8, -1, 9, 10, -1, -1, 11, -1, 12, -1, 13, -1, -1, 14}
	ntree2 := createNaryTree(root2)
	cloneTree(ntree2)
}
