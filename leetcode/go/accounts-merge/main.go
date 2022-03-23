package main

import (
	"fmt"
	"sort"
)

//DFS
var (
	visited  map[string]bool
	adjacent map[string][]string
)

func accountsMergeDFS(accounts [][]string) [][]string {
	visited = make(map[string]bool)
	adjacent = make(map[string][]string)

	for _, account := range accounts {
		accountSize := len(account)
		firstEmail := account[1]

		for j := 2; j < accountSize; j++ {
			email := account[j]
			if _, ok := adjacent[firstEmail]; !ok {
				adjacent[firstEmail] = []string{}
			}
			adjacent[firstEmail] = append(adjacent[firstEmail], email)

			if _, ok := adjacent[email]; !ok {
				adjacent[email] = []string{}
			}
			adjacent[email] = append(adjacent[email], firstEmail)
		}
	}

	var mergedAccounts [][]string
	for _, account := range accounts {
		name := account[0]
		firstEmail := account[1]

		if !visited[firstEmail] {
			var mergedAccount []string
			mergedAccount = append(mergedAccount, name)
			mergedAccount = dfs(mergedAccount, firstEmail)
			sort.Strings(mergedAccount[1:])
			mergedAccounts = append(mergedAccounts, mergedAccount)
		}
	}
	return mergedAccounts
}

func dfs(mergedAccount []string, email string) []string {
	visited[email] = true
	mergedAccount = append(mergedAccount, email)

	if _, ok := adjacent[email]; !ok {
		return mergedAccount
	}
	for _, neighbor := range adjacent[email] {
		if !visited[neighbor] {
			mergedAccount = dfs(mergedAccount, neighbor)
		}
	}
	return mergedAccount
}

//disjoint set
type Disjoint struct {
	representative []int
	size           []int
}

func creatDisjoint(size int) *Disjoint {
	ds := &Disjoint{
		representative: make([]int, size), size: make([]int, size),
	}
	for i := 0; i < size; i++ {
		ds.representative[i] = i
		ds.size[i] = 1
	}
	return ds
}

func (ds *Disjoint) findRepresentative(x int) int {
	if x == ds.representative[x] {
		return x
	}
	ds.representative[x] = ds.findRepresentative(ds.representative[x])
	return ds.representative[x]
}

func (ds *Disjoint) unionBySize(a, b int) {
	repA := ds.findRepresentative(a)
	repB := ds.findRepresentative(b)
	if repA == repB {
		return
	}

	if ds.size[repA] >= ds.size[repB] {
		ds.size[repA] += ds.size[repB]
		ds.representative[repB] = repA
	} else {
		ds.size[repB] += ds.size[repA]
		ds.representative[repA] = repB
	}
}

func accountsMerge2(accounts [][]string) [][]string {
	size := len(accounts)
	ds := creatDisjoint(size)

	emailGroup := make(map[string]int)

	for i := 0; i < size; i++ {
		accountSize := len(accounts[i])
		for j := 1; j < accountSize; j++ {
			email := accounts[i][j]
			if _, ok := emailGroup[email]; !ok {
				emailGroup[email] = i
			} else {
				ds.unionBySize(i, emailGroup[email])
			}
		}
	}

	components := make(map[int][]string)
	for email, val := range emailGroup {
		groupRep := ds.findRepresentative(val)
		if _, ok := components[groupRep]; !ok {
			components[groupRep] = []string{}
		}
		components[groupRep] = append(components[groupRep], email)
	}

	mergedAccounts := [][]string{}
	for group, component := range components {
		sort.Strings(component)
		temp := []string{accounts[group][0]}
		temp = append(temp, component...)
		mergedAccounts = append(mergedAccounts, temp)
	}
	return mergedAccounts
}

//my solution not solved
func accountsMerge(accounts [][]string) [][]string {
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i][0] < accounts[j][0]
	})
	for i := range accounts {
		sort.Strings(accounts[i])
	}
	ret := helper(accounts[1:], accounts[0], [][]string{})
	return ret
}

func helper(accounts [][]string, merged []string, output [][]string) [][]string {
	if len(accounts) == 0 {
		output = appendStringArray(output, merged)
		return output
	}

	var isMerged bool
	first := accounts[0]
	if merged[0] == first[0] {
		for j := 1; j < len(first); j++ {
			if contains(merged, first[j]) {
				isMerged = true
				merged = append(merged, first...)
				merged = deduplicate(merged)
				break
			}
		}
	}

	if isMerged {
		output = helper(accounts[1:], merged, output)
	} else {
		output = appendStringArray(output, merged)
		output = helper(accounts[1:], accounts[0], output)
	}
	return output
}

func appendStringArray(target [][]string, item []string) [][]string {
	temp := make([]string, len(item))
	copy(temp, item)
	temp = deduplicate(temp)
	target = append(target, temp)
	return target
}

func contains(strs []string, item string) bool {
	for _, str := range strs {
		if item == str {
			return true
		}
	}
	return false
}

func deduplicate(s []string) []string {
	m := make(map[string]bool)
	var ret []string
	for _, v := range s {
		if !m[v] {
			m[v] = true
			ret = append(ret, v)
		}
	}
	return ret
}

func main() {
	accounts := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
	fmt.Println(accountsMergeDFS(accounts))
}
