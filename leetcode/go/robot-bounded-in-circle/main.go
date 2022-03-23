package main

import "fmt"

type Robot struct {
	x   int
	y   int
	dir Direction
}

type Direction int

const (
	North Direction = iota
	West
	South
	East
)

func isRobotBounded(instructions string) bool {

	robot := Robot{0, 0, North}
	for i := 0; i < 4; i++ {
		for _, v := range instructions {
			switch string(v) {
			case "G":
				robot.GoRobot()
			case "L":
				robot.TurnLeft()
			case "R":
				robot.TurnRight()
			}
		}
		if robot.x == 0 && robot.y == 0 {
			return true
		}
	}
	return false
}

func (r *Robot) GoRobot() {
	switch r.dir {
	case North:
		r.y -= 1
	case West:
		r.x -= 1
	case South:
		r.y += 1
	case East:
		r.x += 1
	}
}
func (r *Robot) TurnLeft() {
	if r.dir == East {
		r.dir = North
	} else {
		r.dir += 1
	}
}
func (r *Robot) TurnRight() {
	if r.dir == North {
		r.dir = East
	} else {
		r.dir -= 1
	}
}

func main() {
	fmt.Println(isRobotBounded("GGLLGG"))
	fmt.Println(isRobotBounded("GG"))
	fmt.Println(isRobotBounded("GL"))
}
