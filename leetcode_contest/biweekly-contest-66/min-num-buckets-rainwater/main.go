package main

import (
	"fmt"
)

func minimumBuckets(street string) int {
	given := make([]int, len(street))
	H := make([]int, len(street))
	B := make([]int, len(street))

	for i,v := range street{
		if v == 'H' {
			H[i] = 1
		}
	}

	for i,v := range H {
		if v == 1 {
			if i > 0 && H[i-1] == 0{
				B[i-1] += 1
				given[i] += 1
			}
			if i < len(B)-1 && H[i+1] == 0{
				B[i+1] += 1
				given[i] += 1
			}
		}
	}

	for i:=len(given)-1; i>=0; i--{
		if given[i] > 1 {
			if i > 0 && i < len(B)-1 {
				if B[i-1] > B[i+1] {
					B[i+1]--
					given[i]--
				} else {
					B[i-1]--
					given[i]--
				}
			}
		}
	}
	if !equal(given,H) {
		return -1
	}

	var ret int
	for _,v := range B {
		if v>=1 {
			ret++
		}
	}
	return ret
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	street1 := "H..H"
	fmt.Println(minimumBuckets(street1))

	street2 := ".H.H."
	fmt.Println(minimumBuckets(street2))

	street3 := ".HHH."
	fmt.Println(minimumBuckets(street3))

	street4 := "H..H."
	fmt.Println(minimumBuckets(street4))
}
