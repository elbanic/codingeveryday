package main

type Sea struct {
}

func (s *Sea) hasShips(topRight, bottomLeft []int) bool {
	return true
}

func countShips(sea Sea, topRight, bottomLeft []int) int {

	if bottomLeft[0] > topRight[0] || bottomLeft[1] > topRight[1] {
		return 0
	}

	if !sea.hasShips(topRight, bottomLeft) {
		return 0
	}

	if bottomLeft[0] == topRight[0] && bottomLeft[1] == topRight[1] {
		return 1
	}

	midX := (topRight[0] + bottomLeft[0]) / 2
	midY := (topRight[1] + bottomLeft[1]) / 2
	return countShips(sea, []int{midX, midY}, bottomLeft) +
		countShips(sea, topRight, []int{midX + 1, midY + 1}) +
		countShips(sea, []int{midX, topRight[1]}, []int{bottomLeft[0], midY + 1}) +
		countShips(sea, []int{topRight[0], midY}, []int{midX + 1, bottomLeft[1]})
}

func main() {

}
