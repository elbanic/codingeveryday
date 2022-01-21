/*
https://leetcode.com/problems/gas-station/
There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station.
You begin the journey with an empty tank at one of the gas stations.
Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once
in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique

Example 1:
	Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
	Output: 3
	Explanation:
	Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
	Travel to station 4. Your tank = 4 - 1 + 5 = 8
	Travel to station 0. Your tank = 8 - 2 + 1 = 7
	Travel to station 1. Your tank = 7 - 3 + 2 = 6
	Travel to station 2. Your tank = 6 - 4 + 3 = 5
	Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
	Therefore, return 3 as the starting index.

Example 2:
	Input: gas = [2,3,4], cost = [3,4,3]
	Output: -1
	Explanation:
	You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
	Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
	Travel to station 0. Your tank = 4 - 3 + 2 = 3
	Travel to station 1. Your tank = 3 - 3 + 3 = 3
	You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
	Therefore, you can't travel around the circuit once no matter where you start.
*/
package main

import (
	"fmt"
)

func canCompleteCircuit2(gas []int, cost []int) int {
	n := len(gas)
	var totalTank, currTank, start int

	for i := 0; i < n; i++ {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]
		if currTank < 0 {
			start = i + 1
			currTank = 0
		}
	}

	if totalTank >= 0 {
		if start == n {
			return 0
		}
		return start
	}
	return -1
}

func canCompleteCircuit(gas []int, cost []int) int {

	if len(gas) == 0 {
		return -1
	}
	if len(gas) == 1 {
		if gas[0] >= cost[0] {
			return 0
		} else {
			return -1
		}
	}
	for i := 0; i < len(gas); i++ {
		sum := gas[i] - cost[i]
		if sum > 0 {
			cur := runCircuit(i, i+1, sum, gas, cost)
			if cur >= 0 {
				return i
			}
		}
	}
	return -1
}

func runCircuit(start, i, sum int, gas []int, cost []int) int {

	if i >= len(gas) {
		i = 0
	}
	if start == i {
		return sum
	}
	sum += gas[i] - cost[i]
	if sum < 0 {
		return -1
	}
	return runCircuit(start, i+1, sum, gas, cost)
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit2(gas, cost))

	gas2 := []int{2, 3, 4}
	cost2 := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit2(gas2, cost2))
}
