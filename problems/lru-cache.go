package problems

import "fmt"

// ProblemX - LRU Cache
func ProblemX() {

	cache := Construct(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4

	cache.Put(2, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(2)) // returns 2
	cache.Put(1, 1)
	cache.Put(4, 1)
	fmt.Println(cache.Get(2)) // returns -1

}

// LRUCache - ...
type LRUCache struct {
	M [][]int
	C int
}

// Construct - Construct this...
func Construct(capacity int) LRUCache {

	return LRUCache{M: make([][]int, capacity), C: capacity}

}

// Get - Get this...
func (this *LRUCache) Get(key int) int {

	for i, c := len(this.M)-1, 0; i >= 0 && c < this.C; i-- {
		if this.M[i][0] == key {
			this.M = append(this.M, []int{key, this.M[i][1]})
			return this.M[i][1]
		}
		c++
	}
	return -1

}

// Put - Put this...
func (this *LRUCache) Put(key int, value int) {

	this.M = append(this.M, []int{key, value})

	if len(this.M) > this.C {
		this.M = this.M[len(this.M)-this.C:]
	}

}
