/*
https://leetcode.com/problems/design-a-leaderboard/
Design a Leaderboard class, which has 3 functions:
	1. addScore(playerId, score): Update the leaderboard by adding score to the given player's score.
		If there is no player with such id in the leaderboard, add him to the leaderboard with the given score.
	2. top(K): Return the score sum of the top K players.
	3. reset(playerId): Reset the score of the player with the given id to 0 (in other words erase it from the leaderboard).
		It is guaranteed that the player was added to the leaderboard before calling this function.
Initially, the leaderboard is empty.

Example 1:
	Input:
	["Leaderboard","addScore","addScore","addScore","addScore","addScore","top","reset","reset","addScore","top"]
	[[],[1,73],[2,56],[3,39],[4,51],[5,4],[1],[1],[2],[2,51],[3]]
	Output:
	[null,null,null,null,null,null,73,null,null,null,141]

	Explanation:
	Leaderboard leaderboard = new Leaderboard ();
	leaderboard.addScore(1,73);   // leaderboard = [[1,73]];
	leaderboard.addScore(2,56);   // leaderboard = [[1,73],[2,56]];
	leaderboard.addScore(3,39);   // leaderboard = [[1,73],[2,56],[3,39]];
	leaderboard.addScore(4,51);   // leaderboard = [[1,73],[2,56],[3,39],[4,51]];
	leaderboard.addScore(5,4);    // leaderboard = [[1,73],[2,56],[3,39],[4,51],[5,4]];
	leaderboard.top(1);           // returns 73;
	leaderboard.reset(1);         // leaderboard = [[2,56],[3,39],[4,51],[5,4]];
	leaderboard.reset(2);         // leaderboard = [[3,39],[4,51],[5,4]];
	leaderboard.addScore(2,51);   // leaderboard = [[2,51],[3,39],[4,51],[5,4]];
	leaderboard.top(3);           // returns 141 = 51 + 51 + 39;
*/

package main

import (
	"fmt"
	"sort"
)

type player struct {
	id    int
	score int
}
type Leaderboard struct {
	players map[int]int //id, score
	rank    []int
	reCalc  bool
}

func Constructor() Leaderboard {
	return Leaderboard{make(map[int]int), []int{}, true}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	if _, ok := this.players[playerId]; !ok {
		this.players[playerId] = score
	} else {
		this.players[playerId] += score
		score = this.players[playerId]
	}
	if len(this.rank) == 0 {
		this.rank = append(this.rank, playerId)
	}
	this.reCalc = true
}

func (this *Leaderboard) Top(K int) int {
	if this.reCalc {
		var p []player
		for k, v := range this.players {
			p = append(p, player{k, v})
		}
		sort.Slice(p, func(i, j int) bool {
			return p[i].score > p[j].score
		})
		this.rank = make([]int, len(p))
		for i := range p {
			this.rank[i] = p[i].score
		}
		this.reCalc = false
	}
	var sum int
	for i := 0; i < K; i++ {
		sum += this.rank[i]
	}
	return sum
}

func (this *Leaderboard) Reset(playerId int) {
	this.AddScore(playerId, 0)
	delete(this.players, playerId)
}

func main() {
	lboard := Constructor()
	lboard.AddScore(1, 73)
	lboard.AddScore(2, 56)
	lboard.AddScore(3, 39)
	lboard.AddScore(4, 51)
	lboard.AddScore(5, 4)

	fmt.Println(lboard.Top(1))
	lboard.Reset(1)
	lboard.Reset(2)
	lboard.AddScore(2, 51)
	fmt.Println(lboard.Top(3))
}
