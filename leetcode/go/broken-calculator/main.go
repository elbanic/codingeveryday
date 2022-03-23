package main

import (
	"fmt"
	"math"
)

const MaxInt = int(^uint(0) >> 1)

func brokenCalc(startValue int, target int) int {

	var step int

	for target > startValue {
		step++
		if target%2 == 1 {
			target++
		} else {
			target /= 2
		}
	}
	return step + startValue - target
}

func helper(curVal int, target int, step int, minStep int) int {

	if curVal == target {
		return step
	}

	if curVal == 0 {
		return MaxInt
	}

	if minStep < step {
		return minStep
	}

	if curVal > target {
		return helper(curVal-1, target, step+1, minStep)
	} else {
		a := helper(curVal-1, target, step+1, minStep)
		b := helper(curVal*2, target, step+1, minStep)
		return int(math.Min(float64(a), float64(b)))
	}
}

func main() {

	fmt.Println(brokenCalc(2, 3))
	fmt.Println(brokenCalc(5, 8))

}
