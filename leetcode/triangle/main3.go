package main
//
//import "fmt"
//
//var (
//	min    int
//	bFirst bool
//)
//
//func generatePaths(depth int, curIndex int, triangle [][]int, path []int, found chan bool, minTriangle int) int {
//
//	if depth >= len(triangle) {
//		return min
//	}
//	cur := triangle[depth]
//	depth++
//
//	for i, v := range cur {
//		if i == curIndex || i == curIndex+1 {
//			temp := make([]int, len(path)+1)
//			copy(temp, path)
//			temp[len(temp)-1] = v
//
//			if depth == len(triangle) {
//				sum := sum(temp)
//				if !bFirst && sum == minTriangle {
//					found <- true
//				}
//				if bFirst {
//					bFirst = false
//					min = sum
//				} else if !bFirst && sum < min {
//					min = sum
//				}
//			}
//			min = generatePaths(depth, i, triangle, temp, found, minTriangle)
//		}
//	}
//	return min
//}
//
//func minimumTotal(triangle [][]int) int {
//	if len(triangle) == 0 {
//		return 0
//	}
//
//	min = 0
//	bFirst = true
//
//	var found = make(chan bool)
//	var done = make(chan bool)
//	minTriangle := minimumSlices(triangle)
//
//	var ret int
//	go func(triangle [][]int) {
//		ret = generatePaths(0, 0, triangle, nil, found, minTriangle)
//		found <- true
//	}(triangle)
//
//	for {
//		select {
//		case <-found:
//			return min
//		case <-done:
//			return ret
//		}
//	}
//}
//
//func Map(s [][]int, f func([]int) int) []int {
//
//	ss := make([]int, len(s))
//	for i, v := range s {
//		ss[i] = f(v)
//	}
//	return ss
//}
//
//func sum(s []int) int {
//	result := 0
//	for _, v := range s {
//		result += v
//	}
//	return result
//}
//
//func minimumSlices(triangle [][]int) int {
//	var minElementsSum int
//	for _, v := range triangle {
//		minElementsSum += minSlice(v)
//	}
//	return minElementsSum
//}
//
//func minSlice(s []int) int {
//	if len(s) == 0 {
//		return -1
//	}
//	min := s[0]
//	for i, e := range s {
//		if i == 0 || e < min {
//			min = e
//		}
//	}
//	return min
//}
//
//func main() {
//	triangle1 := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
//	fmt.Println(minimumTotal(triangle1))
//}