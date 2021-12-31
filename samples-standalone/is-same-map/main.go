package main

import "fmt"

func isSameMap(a map[string]interface{}, b map[string]interface{}) bool {

	if len(a) != len(b) {
		return false
	}

	var ret bool
	for key, value := range a {
		switch value.(type) {
		case string:
			if b[key].(string) != value.(string) {
				return false
			}
		case map[string]interface{}:
			ret = isSameMap(value.(map[string]interface{}), b[key].(map[string]interface{}))
			if !ret {
				return false
			}
		}
	}
	return true
}

func main() {
	a := make(map[string]interface{})
	b := make(map[string]interface{})

	a["key1"] = "string1"
	a["key2"] = make(map[string]interface{})
	a["key2"].(map[string]interface{})["key3"] = "string2"

	b["key1"] = "string1"
	b["key2"] = make(map[string]interface{})
	b["key2"].(map[string]interface{})["key3"] = "string2"

	fmt.Println(isSameMap(a, b))

	c := make(map[string]interface{})
	d := make(map[string]interface{})

	c["key1"] = "string1"
	c["key2"] = make(map[string]interface{})
	c["key2"].(map[string]interface{})["key3"] = "string2"
	c["key3"] = "string3"

	c["key1"] = "string1"
	c["key2"] = make(map[string]interface{})
	c["key2"].(map[string]interface{})["key3"] = "string2"
	c["key2"].(map[string]interface{})["key4"] = "string3"
	fmt.Println(isSameMap(c, d))
}
