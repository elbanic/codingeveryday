package main

import (
	"fmt"
	"math"
)

type question struct {
	point      int
	brainpower int
}

//iteration
func mostPointsIteration(questions [][]int) int64 {
	type scoreSum struct {
		idx int
		sum int
	}
	var stack []scoreSum
	stack = append(stack, scoreSum{0, 0})

	var max int
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		idx := cur.idx
		sum := cur.sum

		if idx+1 < len(questions)-1 {
			//skip
			stack = append(stack, scoreSum{idx + 1, sum})
		}
		sum += questions[idx][0]
		if idx+questions[idx][1]+1 < len(questions) {
			//solve
			stack = append(stack, scoreSum{idx + questions[idx][1] + 1, sum})
		}
		if max < sum {
			max = sum
		}
	}
	return int64(max)
}

//recursion
func mostPoints(questions [][]int) int64 {
	return int64(helper(0, questions, 0))
}

func helper(idx int, questions [][]int, solvedSum int) int {
	if idx >= len(questions) {
		return solvedSum
	}

	skipMax := helper(idx+1, questions, solvedSum)
	solveMax := helper(idx+questions[idx][1]+1, questions, solvedSum+questions[idx][0])
	return int(math.Max(float64(skipMax), float64(solveMax)))
}

func main() {
	questions := [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}
	fmt.Println(mostPoints(questions))

	questions2 := [][]int{{29, 1}, {90, 5}, {41, 5}, {46, 3}, {49, 5}, {49, 2}, {6, 5}, {36, 5}, {83, 1}, {60, 2}, {97, 3}, {54, 2}, {42, 5}, {42, 2}, {73, 4}, {38, 4}, {16, 4}, {44, 2}, {81, 2}, {76, 3}, {60, 4}, {83, 4}, {94, 2}, {13, 5}, {7, 2}, {77, 2}, {18, 2}, {91, 2}, {43, 4}, {84, 2}, {45, 1}, {42, 5}, {54, 4}, {18, 4}, {96, 5}, {44, 3}, {55, 4}, {49, 3}, {86, 2}, {41, 3}, {54, 3}, {66, 2}, {22, 3}, {35, 5}, {89, 3}, {61, 2}, {1, 3}, {72, 1}, {13, 3}, {70, 4}, {12, 4}, {35, 5}, {16, 3}, {67, 3}, {70, 3}, {5, 4}, {74, 4}, {36, 1}, {47, 2}, {72, 1}}
	fmt.Println(mostPoints(questions2))
}
