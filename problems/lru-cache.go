package problems

import "fmt"

// Problem146 - 146. LRU Cache
func Problem146() {

	// test 0
	//["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	//[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	cache := Construct(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	fmt.Println(cache.Get(4)) // returns 4
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(2)) // returns 2
	fmt.Println(cache.Get(1)) // returns -1
	cache.Put(5, 5)
	fmt.Println(cache.Get(1)) // returns -1
	fmt.Println(cache.Get(2)) // returns 2
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns -1
	fmt.Println(cache.Get(5)) // returns 5

	// test 1
	cache = Construct(1)
	fmt.Println(cache.Get(0)) // returns -1

	// test 2
	cache = Construct(2)

	fmt.Println(cache.Get(2)) // returns -1
	cache.Put(2, 6)
	fmt.Println(cache.Get(1)) // returns -1
	cache.Put(1, 5)
	cache.Put(1, 2)
	fmt.Println(cache.Get(1)) // returns 2
	fmt.Println(cache.Get(2)) // returns 6

	// test 3
	cache = Construct(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4

	// test 4
	cache = Construct(2)

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

	for i := 0; i < len(this.M); i++ {
		if this.M[i] == nil {
			break
		}
		if this.M[i][0] == key {
			ans := this.M[i][1]
			if i > 0 {
				for j := i; j > 0; j-- {
					this.M[j] = this.M[j-1]
				}
				this.M[0] = []int{key, ans}
			}
			return ans
		}
	}
	return -1

}

// Put - Put this...
func (this *LRUCache) Put(key int, value int) {

	if this.M[0] != nil && this.M[0][0] == key {
		this.M[0][1] = value
		return
	}

	i := -1
	for j := 1; j < len(this.M); j++ {
		if this.M[j] == nil || this.M[j][0] == key {
			i = j
			break
		}
	}
	if i > 0 {
		for j := i; j > 0; j-- {
			this.M[j] = this.M[j-1]
		}
	} else {
		this.M = append([][]int{{key, value}}, this.M[:len(this.M)-1]...)
	}
	this.M[0] = []int{key, value}

}
