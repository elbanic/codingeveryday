package main

import (
	"fmt"
	"strconv"
)

func countPoints(rings string) int {

	var rodMaps []map[string]bool
	full := make(map[int]bool)
	for i:=0; i<10; i++ {
		rodMaps = append(rodMaps, make(map[string]bool))
	}

	for i:=0; i<len(rings); i+=2 {
		color := string(rings[i])
		id, err := strconv.Atoi(string(rings[i+1]))
		if err == nil {
			rodMaps[id][color] = true
			if rodMaps[id]["R"] == true && rodMaps[id]["G"] == true && rodMaps[id]["B"] == true {
				full[id] = true
			}
		}
	}
	return len(full)
}

func main() {
	rings := "B0B6G0R6R0R6G9"
	fmt.Println(countPoints(rings))

	rings2 := "G4"
	fmt.Println(countPoints(rings2))
}
