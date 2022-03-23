package main

import (
	"fmt"
	"math"
	"strconv"
)

//brute force
func findMaximumXOR2(nums []int) int {

	var max int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]^nums[j] > max {
				max = nums[i] ^ nums[j]
			}
		}
	}
	return max
}

//bitwise trie
type TrieNode struct {
	children map[string]TrieNode
}

func findMaximumXOR3(nums []int) int {
	maxNum := nums[0]
	for _, num := range nums {
		maxNum = int(math.Max(float64(maxNum), float64(num)))
	}
	L := len(strconv.FormatInt(int64(maxNum), 2))
	n := len(nums)
	bitmask := 1 << L

	strNums := make([]string, n)

	for i := 0; i < n; i++ {
		strNums[i] = strconv.FormatInt(int64(bitmask|nums[i]), 2)[1:]
	}
	trie := TrieNode{make(map[string]TrieNode)}
	var maxXor int

	for _, num := range strNums {
		node := trie
		xorNode := trie
		currXor := 0

		for _, bit := range num {
			if _, ok := node.children[string(bit)]; ok {
				node = node.children[string(bit)]
			} else {
				newNode := TrieNode{make(map[string]TrieNode)}
				node.children[string(bit)] = newNode
				node = newNode
			}

			var toggledBit string
			if string(bit) == "1" {
				toggledBit = "0"
			} else {
				toggledBit = "1"
			}
			if _, ok := xorNode.children[toggledBit]; ok {
				currXor = (currXor << 1) | 1
				xorNode = xorNode.children[toggledBit]
			} else {
				currXor = currXor << 1
				xorNode = xorNode.children[string(bit)]
			}
		}
		maxXor = int(math.Max(float64(maxXor), float64(currXor)))
	}
	return maxXor
}

//hash set
func findMaximumXOR(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	L := len(strconv.FormatInt(int64(max), 2))
	var maxXor, currXor int

	for i := L - 1; i >= 0; i-- {
		maxXor <<= 1
		currXor = maxXor | 1
		prefix := make(map[int]bool)

		for _, v := range nums {
			num := v >> i
			prefix[num] = true
		}
		for p, _ := range prefix {
			if prefix[currXor^p] {
				maxXor = currXor
				break
			}
		}
	}
	return maxXor
}

func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR3(nums))

	nums2 := []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}
	fmt.Println(findMaximumXOR3(nums2))
}
