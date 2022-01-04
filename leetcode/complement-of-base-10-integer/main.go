/*

 */
package main

import (
	"fmt"
	"math"
)

func bitwiseComplement(n int) int {
	bin := toBin(n)
	for i := range bin {
		bin[i] = bin[i] ^ 1
	}

	num := toInt(bin)
	return num
}

func toBin(n int) []int {
	var ret []int
	for n > 1 {
		ret = append(ret, n%2)
		n /= 2
	}
	ret = append(ret, n%2)
	return ret
}

func toInt(n []int) int {
	var ret int
	for i := range n {
		ret += n[i] * int(math.Pow(2, float64(i)))
	}
	return ret
}

func main() {
	fmt.Println(bitwiseComplement(1024))
	fmt.Println(bitwiseComplement2(7))
	fmt.Println(bitwiseComplement(10))
	fmt.Println(bitwiseComplement(5))
}
