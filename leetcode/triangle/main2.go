package main

// func generatePaths(depth int, curIndex int, triangle [][]int, path []int, pathArray [][]int, ch chan []int) [][]int {

// 	if depth >= len(triangle) {
// 		return pathArray
// 	}

// 	cur := triangle[depth]
// 	depth++

// 	for i, v := range cur {
// 		if i == curIndex || i == curIndex+1 {
// 			temp := make([]int, len(path)+1)
// 			copy(temp, path)
// 			temp[len(temp)-1] = v

// 			if depth == len(triangle) {
// 				pathArray = append(pathArray, temp)
// 				ch <- temp
// 			}
// 			pathArray = generatePaths(depth, i, triangle, temp, pathArray, ch)
// 		}
// 	}
// 	return pathArray
// }

// func minimumTotal(triangle [][]int) int {

// 	var ch = make(chan []int, 100)
// 	var done = make(chan bool)

// 	if len(triangle) == 0 {
// 		return 0
// 	}
// 	min := minimumSlices(triangle)

// 	var paths [][]int
// 	go func(triangle [][]int) {
// 		paths = generatePaths(0, 0, triangle, nil, nil, ch)
// 		close(ch)
// 		done <- true
// 	}(triangle)

// L:
// 	for {
// 		select {
// 		case path := <-ch:
// 			if sum(path) == min {
// 				return min
// 			}
// 		case ok := <-done:
// 			if ok {
// 				break L
// 			}
// 		}
// 	}
// 	return minSlice(Map(paths, sum))
// }

// func Map(s [][]int, f func([]int) int) []int {

// 	ss := make([]int, len(s))
// 	for i, v := range s {
// 		ss[i] = f(v)
// 	}
// 	return ss
// }

// func sum(s []int) int {
// 	result := 0
// 	for _, v := range s {
// 		result += v
// 	}
// 	return result
// }

// func minimumSlices(triangle [][]int) int {
// 	var minElementsSum int
// 	for _, v := range triangle {
// 		minElementsSum += minSlice(v)
// 	}
// 	return minElementsSum
// }

// func minSlice(s []int) int {
// 	if len(s) == 0 {
// 		return -1
// 	}
// 	min := s[0]
// 	for i, e := range s {
// 		if i == 0 || e < min {
// 			min = e
// 		}
// 	}
// 	return min
// }

// func main() {
// 	triangle1 := [][]int{{-10}}
// 	fmt.Println(minimumTotal(triangle1))

// 	triangle2 := [][]int{{0}, {0, 0}, {0, 0, 0}}
// 	fmt.Println(minimumTotal(triangle2))
// }
