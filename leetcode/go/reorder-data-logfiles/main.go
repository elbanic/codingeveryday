/*
https://leetcode.com/problems/reorder-data-in-log-files/

You are given an array of logs. Each log is a space-delimited string of words, where the first word is the identifier.
There are two types of logs:
* Letter-logs: All words (except the identifier) consist of lowercase English letters.
* Digit-logs: All words (except the identifier) consist of digits.

Reorder these logs so that:
1. The letter-logs come before all digit-logs.
2. The letter-logs are sorted lexicographically by their contents. If their contents are the same, then sort them lexicographically by their identifiers.
3. The digit-logs maintain their relative ordering.

Return the final order of the logs.

Example 1:
	Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
	Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
	Explanation:
	The letter-log contents are all different, so their ordering is "art can", "art zero", "own kit dig".
	The digit-logs have a relative order of "dig1 8 1 5 1", "dig2 3 6".

Example 2:
	Input: logs = ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
	Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]
 */
package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

type Log struct {
	id string
	cont string
	letter bool
}

func reorderLogFiles(logs []string) []string {
	var letterLogs []Log
	var digitLogs []Log
	for _,str := range logs {
		splits := strings.Split(str, " ")
		cont := splits[1:]
		letter := true
		for _,v := range cont {
			n := new(big.Int)
			if _, err := n.SetString(v, 10); err {
				letter = false
				break
			}
		}
		l := Log{splits[0], strings.Join(cont, " "), letter}
		if letter {
			letterLogs = append(letterLogs, l)
		} else {
			digitLogs = append(digitLogs, l)
		}
	}

	sort.Slice(letterLogs, func(i, j int) bool {
		if letterLogs[i].cont == letterLogs[j].cont {
			return letterLogs[i].id < letterLogs[j].id
		}
		return letterLogs[i].cont < letterLogs[j].cont
	})

	var ret []string
	for i:=0; i<len(letterLogs); i++ {
		ret = append(ret, letterLogs[i].id + " " + letterLogs[i].cont)
	}
	for i:=0; i<len(digitLogs); i++ {
		ret = append(ret, digitLogs[i].id + " " + digitLogs[i].cont)
	}
	return ret
}

func main() {
	logs1 := []string{"dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"}
	fmt.Println(reorderLogFiles(logs1))

	logs2 := []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	fmt.Println(reorderLogFiles(logs2))

}
