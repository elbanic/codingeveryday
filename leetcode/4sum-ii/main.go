package main

import "fmt"

//hash table
func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	twoMap := make(map[int]int)
	var ret int
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			twoMap[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if val, ok := twoMap[-(n3 + n4)]; ok {
				ret += val
			}
		}
	}
	return ret
}

//brute force
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var ret int
	for _, i := range nums1 {
		for _, j := range nums2 {
			for _, k := range nums3 {
				for _, l := range nums4 {
					if i+j+k+l == 0 {
						ret++
						break
					}
				}
			}
		}
	}
	return ret
}

func main() {
	nums1 := []int{-1, -1}
	nums2 := []int{-1, 1}
	nums3 := []int{-1, 1}
	nums4 := []int{1, -1}
	fmt.Println(fourSumCount2(nums1, nums2, nums3, nums4))
}
