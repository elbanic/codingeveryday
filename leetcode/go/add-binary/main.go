/*
https://leetcode.com/problems/add-binary/
*/
package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func addBinary(a string, b string) string {

	aArr := toBinArrayDescend(a)
	bArr := toBinArrayDescend(b)
	resArr := helper(0, 0, aArr, bArr, []int{})
	var res string
	for _, v := range resArr {
		res = strconv.Itoa(v) + res
	}
	return res
}

func toBinArrayDescend(s string) []int {
	var ret []int
	for i := len(s) - 1; i >= 0; i-- {
		val, err := strconv.Atoi(string(s[i]))
		if err == nil {
			ret = append(ret, val)
		}
	}
	return ret
}

func helper(depth, carry int, a []int, b []int, output []int) []int {

	if depth == int(math.Max(float64(len(a)), float64(len(b)))) {
		if carry == 1 {
			output = append(output, carry)
		}
		return output
	}

	sum := carry
	if depth < len(a) {
		sum += a[depth]
	}
	if depth < len(b) {
		sum += b[depth]
	}
	output = append(output, sum%2)
	depth++
	return helper(depth, sum/2, a, b, output)
}

func addBinary2(a string, b string) string {

	var x, y big.Int
	x.SetString(a, 2)
	y.SetString(b, 2)

	for y.Cmp(big.NewInt(0)) != 0 {
		var answer, carry big.Int
		answer.Xor(&x, &y)
		carry.And(&x, &y)
		carry.Lsh(&carry, 1)
		x = answer
		y = carry
	}
	return x.Text(2)
}

func addBinary3(a string, b string) string {
	x, _ := strconv.ParseInt(a, 2, 64)
	y, _ := strconv.ParseInt(b, 2, 64)

	var answer, carry int64
	for y != 0 {
		answer = x ^ y
		carry = (x & y) << 1
		x = answer
		y = carry
	}
	return strconv.FormatInt(x, 2)
}

func main() {
	a := "1111"
	b := "0010"
	fmt.Println(addBinary2(a, b))
	fmt.Println(addBinary3(a, b))
}
