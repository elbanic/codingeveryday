package main

import "fmt"

func goodTriplets(nums1 []int, nums2 []int) int64 {

	type triplet struct {
		x int
		y int
		z int
	}

	tripMap := make(map[triplet]bool)
	for i := 0; i < len(nums1)-2; i++ {
		for j := i + 1; j < len(nums1)-1; j++ {
			for k := j + 1; k < len(nums1); k++ {
				tripMap[triplet{nums1[i], nums1[j], nums1[k]}] = true
			}
		}
	}
	var ret int
	for i := 0; i < len(nums2)-2; i++ {
		for j := i + 1; j < len(nums2)-1; j++ {
			for k := j + 1; k < len(nums2); k++ {
				if tripMap[triplet{nums2[i], nums2[j], nums2[k]}] {
					ret++
				}
			}
		}
	}

	return int64(ret)
}

func main() {
	nums1 := []int{4, 0, 1, 3, 2}
	nums2 := []int{4, 1, 0, 2, 3}
	fmt.Println(goodTriplets(nums1, nums2))
}
