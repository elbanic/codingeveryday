package main

import (
	"fmt"
	"math"
)

func maximumDetonation(bombs [][]int) int {

	if len(bombs) == 0 {
		return 0
	}
	max := 1
	detonationMap := make(map[int][]int)

	for i:=0; i<len(bombs); i++ {
		for j:=0; j<len(bombs); j++ {
			if i!=j {
				if float64(bombs[i][2]) >= distance(bombs[i], bombs[j]) {
					if arr, ok := detonationMap[i]; ok {
						arr = append(arr, j)
						detonationMap[i] = arr
					} else {
						var arr []int
						arr = append(arr, j)
						detonationMap[i] = arr
					}
				}
			}
		}
	}


	for _, val := range detonationMap {
		curDetonate := []int{val[0]}
		var queue []int
		visited := make(map[int]bool)

		for _, v := range val {
			queue = append(queue, v)
		}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if _, ok := visited[cur]; !ok {
				visited[cur] = true
				curDetonate = append(curDetonate, cur)
				for _, v := range detonationMap[cur] {
					queue = append(queue, v)
				}
			}
		}
		if len(curDetonate) >= len(bombs) {
			return len(bombs)
		}
		if len(curDetonate) > max {
			max = len(curDetonate)
		}
	}
	return max
}

func distance(a []int, b []int) float64 {
	ret := math.Sqrt(math.Pow(float64(a[0] - b[0]), 2) + math.Pow(float64(a[1] - b[1]), 2))
	return ret
}

func main() {
	bombs := [][]int{{1,2,3},{2,3,1},{3,4,2},{4,5,3},{5,6,4}}
	fmt.Println(maximumDetonation(bombs))

	bombs1 := [][]int{{4,4,3},{4,4,3}}
	fmt.Println(maximumDetonation(bombs1))
}
