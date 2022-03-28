package main

import (
	"fmt"
	"math"
)

//top-down
func mergeSort(nums []int) []int {
	if len(nums) == 1 || len(nums) == 0 {
		return nums
	}

	l, r := 0, len(nums)
	m := l + (r-l)/2
	left := mergeSort(nums[l:m])
	right := mergeSort(nums[m:r])
	return merge(left, right)
}

func merge(a []int, b []int) []int {
	idxA, idxB := 0, 0
	var ret []int
	for idxA < len(a) || idxB < len(b) {
		if idxA >= len(a) {
			ret = append(ret, b[idxB])
			idxB++
		} else if idxB >= len(b) {
			ret = append(ret, a[idxA])
			idxA++
		} else if a[idxA] < b[idxB] {
			ret = append(ret, a[idxA])
			idxA++
		} else {
			ret = append(ret, b[idxB])
			idxB++
		}
	}
	return ret
}

//bottom-up
func mergeSort2(nums []int) []int {

	if len(nums) == 0 {
		return nil
	}
	for size := 1; size < len(nums); size *= 2 {
		for l := 0; l < len(nums); l += 2 * size {
			m := l + size - 1
			r := int(math.Min(float64(l+2*size-1), float64(len(nums)-1)))
			temp := make([]int, len(nums))
			copy(temp, nums)
			merge2(nums, temp, l, m, r)
		}
	}
	return nums
}

func merge2(arr []int, temp []int, l, m, r int) {
	k, i, j := l, l, m+1
	for i <= m && j <= r {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			j++
		}
	}
	for i <= m {
		temp[k] = arr[i]
		k++
		i++
	}
	for i := l; i <= r; i++ {
		arr[i] = temp[i]
	}
}

func main() {

	nums := []int{1, 5, 3, 2, 8, 7, 6, 4}
	fmt.Println(mergeSort2(nums))
}
