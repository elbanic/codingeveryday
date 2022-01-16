package main

import (
	"fmt"
	"math"
)

//greedy
func minMoves3(target int, maxDoubles int) int {
	depth := 0
	for target > 1 {

		if maxDoubles == 0 {
			return depth + target - 1
		}
		if target%2 == 0 {
			target /= 2
			maxDoubles--
		} else {
			target -= 1
		}
		depth++
	}
	return depth
}

func greedy(depth, target int, maxDoubles int) int {
	if target == 1 {
		return depth
	}
	depth++

	if target%2 == 0 && maxDoubles > 0 {
		target /= 2
		maxDoubles--
	} else {
		target -= 1
	}
	return greedy(depth, target, maxDoubles)
}

//iteration
func minMoves2(target int, maxDoubles int) int {
	type pathInfo struct {
		val   int
		depth int
		maxD  int
	}

	//BFS
	var queue []pathInfo
	queue = append(queue, pathInfo{1, 0, maxDoubles})
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		val := cur.val
		depth := cur.depth
		maxD := cur.maxD

		if val == target {
			return depth
		} else if val > target {
			continue
		}
		depth++
		inc := val + 1
		queue = append(queue, pathInfo{inc, depth, maxD})
		if maxD > 0 {
			maxD--
			doub := cur.val * 2
			queue = append(queue, pathInfo{doub, depth, maxD})
		}
	}
	return -1
}

//recursion
const MAX_INT = int(^uint(0) >> 1)

func minMoves(target int, maxDoubles int) int {
	ret := helper(1, 0, target, maxDoubles)
	return ret
}

func helper(cur, depth, target int, maxDoubles int) int {
	if cur == target {
		return depth
	} else if cur > target {
		return MAX_INT
	}

	depth++
	inc := helper(cur+1, depth, target, maxDoubles)
	doub := MAX_INT
	if maxDoubles > 0 {
		maxDoubles--
		doub = helper(cur*2, depth, target, maxDoubles)
	}
	return int(math.Min(float64(inc), float64(doub)))
}

func main() {
	fmt.Println(minMoves3(5, 0))
	fmt.Println(minMoves3(10, 4))
	fmt.Println(minMoves3(19, 2))
	fmt.Println(minMoves3(656101987, 1))

}
