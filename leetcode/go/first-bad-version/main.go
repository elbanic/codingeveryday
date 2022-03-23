package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	start := 1
	end := n
	var found int

	for start < end {
		cur := start + (end-start)/2
		stat := isBadVersion(cur)
		if stat {
			end = cur
			found = cur
		} else {
			start = cur + 1
		}
	}

	if start == end && isBadVersion(start) {
		return start
	}
	return found
}

func isBadVersion(version int) bool {
	return bad == version
}

var bad int

func setBadVersion(n int) {
	bad = n
}

func main() {
	n := 1
	setBadVersion(1)
	fmt.Println(firstBadVersion(n))
}
