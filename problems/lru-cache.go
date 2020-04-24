package problems

import "fmt"

// ProblemX - LRU Cache
func ProblemX() {

	// test 1
	cache := Construct(1)
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

	return LRUCache{M: make([][]int, 0, capacity), C: capacity}

}

// Get - Get this...
func (this *LRUCache) Get(key int) int {

	for i, c := len(this.M)-1, 0; i >= 0 && c < this.C; i-- {
		if this.M[i][0] == key {
			if c > 0 {
				if i == 0 {
					this.M = append(this.M[1:], []int{key, this.M[i][1]})
				} else {
					this.M = append(this.M[:i], this.M[i+1:]...)
					this.M = append(this.M, []int{key, this.M[i][1]})
				}
			}
			return this.M[len(this.M)-1][1]
		}
		c++
	}
	return -1

}

// Put - Put this...
func (this *LRUCache) Put(key int, value int) {

	if len(this.M) == 0 {
		this.M = append(this.M, []int{key, value})
	} else {
		for i, c := len(this.M)-1, 0; i >= 0 && c < this.C; i-- {
			if this.M[i][0] == key {
				if c == 0 {
					this.M[i][1] = value
				} else {
					if i == 0 {
						this.M = append(this.M[1:], []int{key, value})
					} else {
						this.M = append(this.M[:i], this.M[i+1:]...)
						this.M = append(this.M, []int{key, value})
					}
				}
				return
			}
			c++
		}
		if len(this.M) < this.C {
			this.M = append(this.M, []int{key, value})
		} else {
			this.M = append(this.M[1:], []int{key, value})
		}
	}

}
