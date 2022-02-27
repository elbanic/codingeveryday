package main

import "fmt"

func minSteps(s string, t string) int {

	sMap := make(map[string]int)
	for _, i := range s {
		sMap[string(i)]++
	}
	tMap := make(map[string]int)
	for _, i := range t {
		tMap[string(i)]++
	}
	var cnt int
	for k, sVal := range sMap {
		if tVal, ok := tMap[k]; ok {
			if tVal > sVal {
				cnt += tVal - sVal
			}
		} else {
			cnt += sVal
		}
	}
	for k, tVal := range tMap {
		if sVal, ok := sMap[k]; ok {
			if tVal < sVal {
				cnt += sVal - tVal
			}
		} else {
			cnt += tVal
		}
	}

	return cnt
}

func main() {
	s := "cotxazilut"
	t := "nahrrmcchxwrieqqdwdpneitkxgnt"
	fmt.Println(minSteps(s, t))
}
