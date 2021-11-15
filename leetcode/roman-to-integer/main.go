/*
https://leetcode.com/problems/roman-to-integer/


Roman numerals are represented by seven different symbols:I,V,X,L,C,D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, `2` is written as `II` in Roman numeral, just two one's added together. `12` is written as `XII`, which is simply `X + II`. The number `27` is written as `XXVII`, which is `XX + V + II`.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not `IIII`. Instead, the number four is written as `IV`. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as `IX`. There are six instances where subtraction is used:

- `I` can be placed before `V`(5) and `X`(10) to make 4 and 9.
- `X` can be placed before `L`(50) and `C`(100) to make 40 and 90.
- `C` can be placed before `D`(500) and `M`(1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.

Example 1:
	Input: s = "III"
	Output: 3

Example 2:
	Input: s = "IV"
	Output: 4

Example 3:
	Input: s = "IX"
	Output: 9

Example 4:
	Input: s = "LVIII"
	Output: 58
	Explanation: L = 50, V= 5, III = 3.

Example 5:
	Input: s = "MCMXCIV"
	Output: 1994
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 */

package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	romanMap := make(map[string]int)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000

	var ret int
	stringSlice := strings.Split(s,"")
	for i:=0; i<len(stringSlice); i++ {
		if (stringSlice[i] == "I" || stringSlice[i] == "X" || stringSlice[i] == "C") &&
			i != len(stringSlice)-1 {
			if stringSlice[i] == "I" && (stringSlice[i+1]=="V" ||stringSlice[i+1]=="X") {
				ret -= 1
			} else if stringSlice[i] == "X" && (stringSlice[i+1]=="L" ||stringSlice[i+1]=="C") {
				ret -= 10
			} else if stringSlice[i] == "C" && (stringSlice[i+1]=="D" ||stringSlice[i+1]=="M") {
				ret -= 100
			} else {
				ret += romanMap[stringSlice[i]]
			}
		} else {
			ret += romanMap[stringSlice[i]]
		}
	}
	return ret
}

func main() {

	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))

}
