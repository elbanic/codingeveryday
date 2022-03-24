package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {

	var boats [][]int

	sort.Slice(people, func(i, j int) bool {
		return people[i] > people[j]
	})

	for len(people) > 0 {
		cur := people[0]
		group := []int{cur}

		capacity := limit - cur
		people = people[1:]

		if len(people) > 0 && capacity > 0 {
			for i, p := range people {
				if p <= capacity {
					group = append(group, p)
					if i+1 > len(people) {
						people = []int{}
					} else {
						people = append(people[:i], people[i+1:]...)
					}
					capacity -= p
				}
				if capacity == 0 {
					break
				}
			}
		}
		boats = append(boats, group)
		capacity = limit
		group = []int{}
	}

	return len(boats)
}

func main() {
	people := []int{3, 2, 3, 2, 2}
	limit := 6
	fmt.Println(numRescueBoats(people, limit))
}
