package main

import (
	"fmt"
	"sync"
)

func addSpaces(s string, spaces []int) string {

	output := new(string)

	div := len(spaces) / 8
	if div >= 10 {
		var sSplit []string
		var sCur int
		var spacesSplit [][]int
		var i int
		for i < len(spaces)-div {
			temp := make([]int, len(spaces[i:i+div]))
			copy(temp, spaces[i:i+div])
			for j := range temp {
				temp[j] = temp[j] - sCur
			}
			sSplit = append(sSplit, s[sCur:spaces[i+div-1]])
			spacesSplit = append(spacesSplit, temp)
			sCur = spaces[i+div-1]
			i += div
		}
		temp := make([]int, len(spaces[i:]))
		copy(temp, spaces[i:])
		for j := range temp {
			temp[j] = temp[j] - sCur
		}
		sSplit = append(sSplit, s[sCur:])
		spacesSplit = append(spacesSplit, temp)

		outputs := make([]string, len(spacesSplit))
		var wg sync.WaitGroup
		for i := range spacesSplit {
			wg.Add(1)
			go func(s string, spaces []int, output *string) {
				addSpacesHelper(s, spaces, output)
				wg.Done()
			}(sSplit[i], spacesSplit[i], &outputs[i])
		}
		wg.Wait()

		var ret string
		for _, v := range outputs {
			ret += v
		}
		return ret
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			addSpacesHelper(s, spaces, output)
			wg.Done()
		}()
		wg.Wait()

		return *output
	}
}

func addSpacesHelper(s string, spaces []int, output *string) {
	var loc int
	for _, v := range spaces {
		*output += s[loc:v] + " "
		loc = v
	}
	*output += s[loc:]
}

func main() {
	s := "spacing"
	spaces := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(addSpaces(s, spaces))
}
