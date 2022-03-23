package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 {
		return false
	}

	var curr, prev, next int
	for i := 0; i < len(flowerbed)-1; i++ {
		curr = flowerbed[i]
		next = flowerbed[i+1]
		if prev == 0 && curr == 0 && next == 0 {
			n--
			flowerbed[i] = 1
		}
		if n <= 0 {
			return true
		}
		prev = flowerbed[i]
	}
	next = flowerbed[len(flowerbed)-1]
	if prev == 0 && curr == 0 && next == 0 {
		n--
	}
	if n <= 0 {
		return true
	}
	return false
}

func main() {
	flowerbed := []int{0}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))

	flowerbed2 := []int{1, 0, 0, 0, 1}
	n2 := 2
	fmt.Println(canPlaceFlowers(flowerbed2, n2))
}
