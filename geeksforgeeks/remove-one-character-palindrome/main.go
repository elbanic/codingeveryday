/*
https://www.geeksforgeeks.org/remove-character-string-make-palindrome/
Given a string, determine whether it is a valid palindrome or whether
it can be a valid palindrome after removing one character.

Example 1:
	“aba” → true

Example 2:
	“abca” → true

Example 3:
	"abcdba" → true
 */
package main

import "fmt"

func palindromeOneRemove(str string) bool {

	var arr []string
	for _, v := range str {
		arr = append(arr, string(v))
	}

	ret := true
	removeIndex1, removeIndex2 := -1, -1

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			removeIndex1 = i
			removeIndex2 = len(arr) - 1 - i
			ret = false
			break
		}
	}
	if ret {
		return ret
	}

	arr1 := make([]string, len(arr))
	copy(arr1, arr)
	arr1 = append(arr1[:removeIndex1], arr1[removeIndex1+1:]...)

	ret = true
	for i := 0; i < len(arr1)/2; i++ {
		if arr1[i] != arr1[len(arr1)-1-i] {
			ret = false
			break
		}
	}
	if ret {
		return ret
	}

	arr2 := make([]string, len(arr))
	copy(arr2, arr)
	arr2 = append(arr[:removeIndex2], arr[removeIndex2+1:]...)

	ret = true
	for i := 0; i < len(arr2)/2; i++ {
		if arr2[i] != arr2[len(arr2)-1-i] {
			ret = false
			break
		}
	}
	return ret
}

func main() {
	fmt.Println(palindromeOneRemove("abcba"))
	fmt.Println(palindromeOneRemove("abcbea"))
	fmt.Println(palindromeOneRemove("abecbea"))
}
