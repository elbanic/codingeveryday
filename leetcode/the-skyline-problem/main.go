package main

import "fmt"

type building struct {
	l int
	r int
	h int
}

type point struct {
	x int
	y int
}

var R int

func getSkyline(buildings [][]int) [][]int {
	R = 0
	var bdArr []building
	removeDup := make(map[building]bool)
	for i := range buildings {
		if i == len(buildings)-1 {
			R = buildings[i][1]
		}

		newBD := building{buildings[i][0], buildings[i][1], buildings[i][2]}
		if removeDup[newBD] {
			continue
		}
		removeDup[newBD] = true
		bdArr = append(bdArr, building{buildings[i][0], buildings[i][1], buildings[i][2]})
	}
	return helperHorizon(bdArr, point{-1, 0}, point{R, 0}, -1, [][]int{})
}

func helperHorizon(buildings []building, ray1, ray2 point, curBDId int, output [][]int) [][]int {

	if ray1.x == ray2.x && ray2.x >= R {
		return output
	}

	var found bool
	minClose, maxBDId := -1, -1
	var intersectPT point
	for i := range buildings {
		if curBDId != -1 {
			if i == curBDId {
				continue
			}
		}
		if minClose != -1 {
			if minClose < buildings[i].l {
				continue
			}
		}

		B1, B2 := point{buildings[i].l, 0}, point{buildings[i].l, buildings[i].h}
		pt, ok := intersectPoint(ray1, ray2, B1, B2)
		if ok {
			if !found {
				maxBDId = i
				minClose = buildings[i].l
				intersectPT = pt
				found = true
			} else if buildings[maxBDId].h <= buildings[i].h {
				if buildings[maxBDId].h == buildings[i].h {
					if buildings[maxBDId].r < buildings[i].r {
						maxBDId = i
						intersectPT = pt
					}
				} else {
					maxBDId = i
					intersectPT = pt
				}
			}
		}
	}
	if found {
		leftTop := []int{intersectPT.x, buildings[maxBDId].h}
		if ray1.y != leftTop[1] || leftTop[0] != buildings[maxBDId].l || leftTop[1] != buildings[maxBDId].h {
			output = append(output, leftTop)
		}
		output = helperHorizon(buildings, point{leftTop[0], leftTop[1]}, point{buildings[maxBDId].r, buildings[maxBDId].h}, maxBDId, output)
	} else {
		output = helperVertical(buildings, point{buildings[curBDId].r, buildings[curBDId].h}, point{buildings[curBDId].r, 0}, curBDId, output)
	}
	return output
}

func helperVertical(buildings []building, ray1, ray2 point, curBDId int, output [][]int) [][]int {
	var found bool
	minClose, maxBDId := -1, -1
	var intersectPT point
	for i := range buildings {
		if curBDId != -1 {
			if i == curBDId {
				continue
			}
		}
		if minClose != -1 {
			if minClose > buildings[i].h {
				continue
			}
		}

		B1, B2 := point{buildings[i].l, buildings[i].h}, point{buildings[i].r, buildings[i].h}
		pt, ok := intersectPoint(ray1, ray2, B1, B2)
		if ok {
			if !found {
				maxBDId = i
				minClose = buildings[i].h
				intersectPT = pt
				if pt.x < buildings[i].r {
					found = true
				}
			} else if buildings[maxBDId].h <= buildings[i].h {
				if buildings[maxBDId].h == buildings[i].h {
					if buildings[maxBDId].r < buildings[i].r {
						maxBDId = i
						intersectPT = pt
					}
				} else {
					maxBDId = i
					intersectPT = pt
				}
			}
		}
	}
	if found {
		leftTop := []int{intersectPT.x, intersectPT.y}
		output = append(output, leftTop)
		output = helperHorizon(buildings, point{leftTop[0], leftTop[1]}, point{buildings[maxBDId].r, buildings[maxBDId].h}, maxBDId, output)
	} else {
		floor := []int{buildings[curBDId].r, 0}
		output = append(output, floor)
		output = helperHorizon(buildings, point{floor[0], 0}, point{R, 0}, curBDId, output)
	}
	return output
}

func intersectPoint(A1 point, A2 point, B1 point, B2 point) (point, bool) {

	var pt point
	var t, s float64
	divider := (B2.y-B1.y)*(A2.x-A1.x) - (B2.x-B1.x)*(A2.y-A1.y)
	if divider == 0 {
		return pt, false
	}

	_t := (B2.x-B1.x)*(A1.y-B1.y) - (B2.y-B1.y)*(A1.x-B1.x)
	_s := (A2.x-A1.x)*(A1.y-B1.y) - (A2.y-A1.y)*(A1.x-B1.x)
	if _t == 0 && _s == 0 {
		return pt, false
	}

	t = float64(_t) / float64(divider)
	s = float64(_s) / float64(divider)
	if t < 0 || t > 1 || s < 0 || s > 1 {
		return pt, false
	}

	pt.x = A1.x + int(t*float64(A2.x-A1.x))
	pt.y = A1.y + int(t*float64(A2.y-A1.y))
	return pt, true
}

func main() {

	//buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	buildings := [][]int{{1, 5, 3}, {1, 5, 3}, {1, 5, 3}}
	fmt.Println(getSkyline(buildings))
}
