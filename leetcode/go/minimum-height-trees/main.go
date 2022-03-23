/*
https://leetcode.com/problems/minimum-height-trees/
A tree is an undirected graph in which any two vertices are connected by exactly one path.
In other words, any connected graph without simple cycles is a tree.
Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] = [ai, bi] indicates
that there is an undirected edge between the two nodes ai and bi in the tree, you can choose any node of the tree as the root.
When you select a node x as the root, the result tree has height h. Among all possible rooted trees,
those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).

Return a list of all MHTs' root labels. You can return the answer in any order.
The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.

Example 1:
	Input: n = 4, edges = [[1,0],[1,2],[1,3]]
	Output: [1]
	Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is the only MHT.

Example 2:
	Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
	Output: [3,4]

Example 3:
	Input: n = 1, edges = []
	Output: [0]

Example 4:
	Input: n = 2, edges = [[0,1]]
	Output: [0,1]
*/

package main

import (
	"fmt"
)

type labelHeight struct {
	l int
	h int
}

//centroid
func findMinHeightTrees2(n int, edges [][]int) []int {
	if n < 2 {
		var centroids []int
		for i := 0; i < n; i++ {
			centroids = append(centroids, i)
		}
		return centroids
	}
	var neighbors []map[int]bool
	for i := 0; i < n; i++ {
		neighbors = append(neighbors, make(map[int]bool))
	}

	for _, edge := range edges {
		start := edge[0]
		end := edge[1]
		neighbors[start][end] = true
		neighbors[end][start] = true
	}

	var leaves []int
	for i := 0; i < n; i++ {
		if len(neighbors[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	remainingNodes := n
	for remainingNodes > 2 {
		remainingNodes -= len(leaves)
		var newLeaves []int

		for _, leaf := range leaves {
			var neighbor int
			for key := range neighbors[leaf] {
				neighbor = key
				break
			}
			delete(neighbors[neighbor], leaf)
			if len(neighbors[neighbor]) == 1 {
				newLeaves = append(newLeaves, neighbor)
			}
		}
		leaves = newLeaves
	}
	return leaves
}

func findMinHeightTrees(n int, edges [][]int) []int {

	var minDepth int
	heights := make(map[int]int, n)
	for i := 0; i < n; i++ {
		depth := 0
		queue := []labelHeight{labelHeight{i, 0}}
		visited := make(map[int]bool)
		for len(queue) > 0 {

			cur := queue[0]
			queue = queue[1:]
			visited[cur.l] = true
			if i != 0 && cur.h > minDepth {
				depth = cur.h
				break
			}

			for _, v := range edges {
				if !visited[v[1]] && v[0] == cur.l {
					queue = append(queue, labelHeight{v[1], cur.h + 1})
				}
				if !visited[v[0]] && v[1] == cur.l {
					queue = append(queue, labelHeight{v[0], cur.h + 1})
				}
			}

			if cur.h > depth {
				depth = cur.h
			}
		}
		if i == 0 {
			minDepth = depth
		} else {
			if minDepth > depth {
				minDepth = depth
			}
		}
		heights[i] = depth
	}

	var ret []int
	for k, v := range heights {
		if v == minDepth {
			ret = append(ret, k)
		}
	}
	return ret
}

func main() {
	n := 5
	edges := [][]int{{1, 0}, {1, 2}, {1, 3}, {4, 3}}
	fmt.Println(findMinHeightTrees2(n, edges))

	n1 := 378
	edges1 := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 4}, {1, 5}, {0, 6}, {6, 7}, {2, 8}, {5, 9}, {0, 10}, {3, 11}, {3, 12}, {10, 13}, {9, 14}, {4, 15}, {14, 16}, {5, 17}, {10, 18}, {3, 19}, {13, 20}, {6, 21}, {6, 22}, {8, 23}, {20, 24}, {10, 25}, {14, 26}, {23, 27}, {10, 28}, {25, 29}, {28, 30}, {25, 31}, {9, 32}, {18, 33}, {4, 34}, {13, 35}, {11, 36}, {34, 37}, {20, 38}, {25, 39}, {19, 40}, {9, 41}, {9, 42}, {22, 43}, {38, 44}, {1, 45}, {5, 46}, {5, 47}, {13, 48}, {30, 49}, {16, 50}, {10, 51}, {22, 52}, {26, 53}, {39, 54}, {16, 55}, {45, 56}, {3, 57}, {42, 58}, {29, 59}, {50, 60}, {29, 61}, {51, 62}, {0, 63}, {24, 64}, {43, 65}, {57, 66}, {2, 67}, {36, 68}, {58, 69}, {16, 70}, {52, 71}, {16, 72}, {19, 73}, {34, 74}, {20, 75}, {59, 76}, {9, 77}, {43, 78}, {26, 79}, {47, 80}, {58, 81}, {10, 82}, {43, 83}, {79, 84}, {64, 85}, {59, 86}, {6, 87}, {51, 88}, {65, 89}, {71, 90}, {18, 91}, {2, 92}, {27, 93}, {19, 94}, {27, 95}, {50, 96}, {81, 97}, {11, 98}, {18, 99}, {58, 100}, {32, 101}, {67, 102}, {56, 103}, {29, 104}, {53, 105}, {45, 106}, {104, 107}, {7, 108}, {80, 109}, {22, 110}, {33, 111}, {30, 112}, {99, 113}, {41, 114}, {75, 115}, {107, 116}, {90, 117}, {6, 118}, {83, 119}, {27, 120}, {0, 121}, {48, 122}, {72, 123}, {62, 124}, {99, 125}, {67, 126}, {69, 127}, {15, 128}, {110, 129}, {120, 130}, {127, 131}, {73, 132}, {102, 133}, {44, 134}, {130, 135}, {69, 136}, {72, 137}, {110, 138}, {135, 139}, {88, 140}, {80, 141}, {116, 142}, {26, 143}, {14, 144}, {74, 145}, {9, 146}, {127, 147}, {110, 148}, {120, 149}, {113, 150}, {57, 151}, {17, 152}, {42, 153}, {60, 154}, {26, 155}, {102, 156}, {33, 157}, {70, 158}, {13, 159}, {78, 160}, {41, 161}, {16, 162}, {37, 163}, {69, 164}, {131, 165}, {59, 166}, {78, 167}, {69, 168}, {134, 169}, {109, 170}, {122, 171}, {96, 172}, {46, 173}, {30, 174}, {62, 175}, {114, 176}, {111, 177}, {82, 178}, {169, 179}, {177, 180}, {170, 181}, {119, 182}, {138, 183}, {8, 184}, {10, 185}, {126, 186}, {9, 187}, {50, 188}, {36, 189}, {49, 190}, {34, 191}, {72, 192}, {80, 193}, {167, 194}, {193, 195}, {79, 196}, {179, 197}, {95, 198}, {48, 199}, {57, 200}, {17, 201}, {76, 202}, {158, 203}, {78, 204}, {153, 205}, {137, 206}, {9, 207}, {206, 208}, {71, 209}, {62, 210}, {162, 211}, {154, 212}, {4, 213}, {126, 214}, {198, 215}, {149, 216}, {162, 217}, {29, 218}, {24, 219}, {18, 220}, {114, 221}, {170, 222}, {118, 223}, {36, 224}, {201, 225}, {185, 226}, {152, 227}, {193, 228}, {176, 229}, {135, 230}, {118, 231}, {68, 232}, {14, 233}, {79, 234}, {144, 235}, {226, 236}, {171, 237}, {172, 238}, {162, 239}, {72, 240}, {133, 241}, {123, 242}, {54, 243}, {91, 244}, {90, 245}, {219, 246}, {203, 247}, {156, 248}, {149, 249}, {53, 250}, {235, 251}, {83, 252}, {11, 253}, {92, 254}, {191, 255}, {60, 256}, {206, 257}, {4, 258}, {233, 259}, {177, 260}, {176, 261}, {17, 262}, {147, 263}, {155, 264}, {27, 265}, {249, 266}, {144, 267}, {44, 268}, {125, 269}, {42, 270}, {39, 271}, {50, 272}, {64, 273}, {15, 274}, {265, 275}, {75, 276}, {104, 277}, {239, 278}, {108, 279}, {64, 280}, {277, 281}, {179, 282}, {183, 283}, {131, 284}, {129, 285}, {16, 286}, {89, 287}, {149, 288}, {57, 289}, {230, 290}, {34, 291}, {266, 292}, {22, 293}, {41, 294}, {288, 295}, {194, 296}, {237, 297}, {104, 298}, {218, 299}, {152, 300}, {138, 301}, {112, 302}, {62, 303}, {112, 304}, {186, 305}, {37, 306}, {163, 307}, {158, 308}, {196, 309}, {3, 310}, {264, 311}, {294, 312}, {232, 313}, {91, 314}, {56, 315}, {204, 316}, {299, 317}, {221, 318}, {118, 319}, {82, 320}, {295, 321}, {316, 322}, {219, 323}, {124, 324}, {226, 325}, {230, 326}, {316, 327}, {12, 328}, {29, 329}, {114, 330}, {247, 331}, {206, 332}, {183, 333}, {74, 334}, {235, 335}, {280, 336}, {109, 337}, {241, 338}, {245, 339}, {92, 340}, {236, 341}, {210, 342}, {269, 343}, {194, 344}, {306, 345}, {139, 346}, {232, 347}, {313, 348}, {19, 349}, {32, 350}, {39, 351}, {289, 352}, {119, 353}, {306, 354}, {281, 355}, {239, 356}, {235, 357}, {106, 358}, {273, 359}, {334, 360}, {68, 361}, {199, 362}, {215, 363}, {207, 364}, {90, 365}, {49, 366}, {10, 367}, {348, 368}, {47, 369}, {239, 370}, {271, 371}, {349, 372}, {167, 373}, {96, 374}, {48, 375}, {23, 376}, {256, 377}}
	fmt.Println(findMinHeightTrees2(n1, edges1))
}
