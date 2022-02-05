package main

import (
	"fmt"
	"strconv"
	"strings"
)

//brute force
func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {

	poss := getPossibility(targetSeconds)
	minCost := int(^uint(0) >> 1)

	for _, possible := range poss {
		var cost int
		s := strings.TrimLeft(possible.Min+possible.Sec, "0")
		prevFinger := strconv.Itoa(startAt)
		for i := 0; i < len(s); i++ {
			if string(s[i]) != prevFinger {
				cost += moveCost
			}
			cost += pushCost
			prevFinger = string(s[i])
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

type MinSec struct {
	Min string
	Sec string
}

func getPossibility(targetSec int) []MinSec {

	var ret []MinSec
	cur := targetSec % 60
	min := targetSec / 60
	if min < 100 {
		ret = append(ret, MinSec{strconv.Itoa(min), fmt.Sprintf("%02d", cur)})
	}

	if cur < 40 && min > 0 {
		cur += 60
		min -= 1
		ret = append(ret, MinSec{strconv.Itoa(min), fmt.Sprintf("%02d", cur)})
	}
	return ret
}

func main() {
	fmt.Println(minCostSetTime(1, 9403, 9402, 6008))

	fmt.Println(minCostSetTime(0, 1, 2, 76))
}
