/*
https://leetcode.com/problems/two-sum/
Given an array of integers `nums` and an integer `target`, return *indices of the two numbers such that they add up to `target`*.
You may assume that each input would have ***exactly* one solution**, and you may not use the *same* element twice.
You can return the answer in any order.

Example 1:

	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

	Input: nums = [3,2,4], target = 6
	Output: [1,2]

Example 3:

	Input: nums = [3,3], target = 6
	Output: [0,1]

Constraints:

	2 <= nums.length <= 104
	-109 <= nums[i] <= 109
	-109 <= target <= 109
	Only one valid answer exists.

Approach

	I will use a nested loop combined by two loops.
	A loop is one item in the array and B loop is one item in the array except for the item of A loop.
	If sum of the two items are the same with target, then the pair will be stored in a channel.
*/

package main

import (
	"fmt"
	"sync"
)

func twoSum(nums []int, target int) []int {
	ch := make(chan []int)
	defer close(ch)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(nums []int) {
		wg.Done()
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					v := []int{i, j}
					ch <- v
				}
			}
		}
	}(nums)
	wg.Wait()

	results := <-ch
	return results
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
