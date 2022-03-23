/*
https://leetcode.com/problems/lru-cache/
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
Implement the LRUCache class:
	* LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
	* int get(int key) Return the value of the key if the key exists, otherwise return -1.
	* void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache.
		If the number of keys exceeds the capacity from this operation, evict the least recently used key.
	*  The functions get and put must each run in O(1) average time complexity.

Example 1:
	Input
	["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	Output
	[null, null, null, 1, null, -1, null, -1, 3, 4]

	Explanation
	LRUCache lRUCache = new LRUCache(2);
	lRUCache.put(1, 1); // cache is {1=1}
	lRUCache.put(2, 2); // cache is {1=1, 2=2}
	lRUCache.get(1);    // return 1
	lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	lRUCache.get(2);    // returns -1 (not found)
	lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	lRUCache.get(1);    // return -1 (not found)
	lRUCache.get(3);    // return 3
	lRUCache.get(4);    // return 4
*/
package main

import (
	"fmt"
	"time"
)

type LRUCache struct {
	cap     int
	keyVale map[int]int
	freq    map[int]time.Time
}

func Constructor(capacity int) LRUCache {
	m1 := make(map[int]int)
	m2 := make(map[int]time.Time)
	return LRUCache{capacity, m1, m2}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.keyVale[key]; ok {
		this.freq[key] = time.Now()
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.freq[key] = time.Now()
	if _, ok := this.keyVale[key]; ok {
		this.keyVale[key] = value
	} else {
		if len(this.keyVale)+1 > this.cap {
			this.deleteLRU()
		}
		this.keyVale[key] = value
	}
}

func (this *LRUCache) deleteLRU() {
	min := time.Now()
	key := -1
	for k, date := range this.freq {
		if date.Before(min) {
			min = date
			key = k
		}
	}
	if key != -1 {
		delete(this.keyVale, key)
		delete(this.freq, key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))

	lru2 := Constructor(2)
	lru2.Put(2, 1)
	lru2.Put(1, 1)
	lru2.Put(2, 3)
	lru2.Put(4, 1)
	fmt.Println(lru2.Get(1))
	fmt.Println(lru2.Get(2))

}
