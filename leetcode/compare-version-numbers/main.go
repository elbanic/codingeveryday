package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {

	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	N := int(math.Max(float64(len(v1)), float64(len(v2))))

	for i := 0; i < N; i++ {
		var curV1, curV2 int
		if i >= len(v1) {
			curV1 = 0
		} else {
			curV1, _ = strconv.Atoi(v1[i])
		}
		if i >= len(v2) {
			curV2 = 0
		} else {
			curV2, _ = strconv.Atoi(v2[i])
		}
		if curV1 > curV2 {
			return 1
		} else if curV1 < curV2 {
			return -1
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
}
