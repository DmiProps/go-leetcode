package problems

import "fmt"

// ProblemX - First Unique Number
func ProblemX() {

	// test 1
	firstUnique := Constructor2([]int{2, 3, 5})
	fmt.Println(firstUnique.ShowFirstUnique()) // return 2
	firstUnique.Add(5)                         // the queue is now [2,3,5,5]
	fmt.Println(firstUnique.ShowFirstUnique()) // return 2
	firstUnique.Add(2)                         // the queue is now [2,3,5,5,2]
	fmt.Println(firstUnique.ShowFirstUnique()) // return 3
	firstUnique.Add(3)                         // the queue is now [2,3,5,5,2,3]
	fmt.Println(firstUnique.ShowFirstUnique()) // return -1

	// test 2
	/*irstUnique firstUnique = new FirstUnique([7,7,7,7,7,7]);
	firstUnique.showFirstUnique(); // return -1
	firstUnique.add(7);            // the queue is now [7,7,7,7,7,7,7]
	firstUnique.add(3);            // the queue is now [7,7,7,7,7,7,7,3]
	firstUnique.add(3);            // the queue is now [7,7,7,7,7,7,7,3,3]
	firstUnique.add(7);            // the queue is now [7,7,7,7,7,7,7,3,3,7]
	firstUnique.add(17);           // the queue is now [7,7,7,7,7,7,7,3,3,7,17]
	firstUnique.showFirstUnique(); // return 17

	// test 3
	FirstUnique firstUnique = new FirstUnique([809]);
	firstUnique.showFirstUnique(); // return 809
	firstUnique.add(809);          // the queue is now [809,809]
	firstUnique.showFirstUnique(); // return -1*/

}

// FirstUnique - this struct
type FirstUnique struct {
	m map[int]int
	l int
}

// Constructor2 - cnstructor
func Constructor2(nums []int) FirstUnique {
	m := make(map[int]int)
	l := len(nums)
	for idx, n := range nums {
		if _, ok := m[n]; ok {
			m[n] = -1
		} else {
			m[n] = idx
		}
	}
	return FirstUnique{m: m, l: l}
}

// ShowFirstUnique - this show
func (this *FirstUnique) ShowFirstUnique() int {

	min := -1
	ans := -1
	for n, idx := range this.m {
		if idx >= 0 {
			if min < 0 || idx < min {
				min = idx
				ans = n
			}
		}
	}
	return ans

}

// Add - this add value
func (this *FirstUnique) Add(value int) {
	if _, ok := this.m[value]; ok {
		this.m[value] = -1
	} else {
		this.m[value] = this.l
		this.l++
	}
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
