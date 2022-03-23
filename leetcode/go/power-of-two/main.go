package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo2(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	return n == 1
}

func isPowerOfTwo(n int) bool {
	for i := 0; i < n; i++ {
		if int(math.Pow(2, float64(i))) == n {
			return true
		}
		if int(math.Pow(2, float64(i))) > n {
			break
		}
	}
	return false
}

//func isPowerOfTwo(n int) bool {
//
//	if n == 1 {
//		return true
//	}
//
//	i := n / 2
//	left := 0
//	right := n
//	for left < right {
//		if left+1 == right {
//			return false
//		}
//		if i*i == n {
//			return true
//		}
//		if i*i > n {
//			right = i
//			i = i + (left-i)/2
//		} else {
//			left = i
//			i = i + (right-i)/2
//		}
//	}
//	return false
//}

func main() {
	fmt.Println(isPowerOfTwo(int(math.Pow(2, 2))))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(100))
	fmt.Println(isPowerOfTwo(1024))
}
