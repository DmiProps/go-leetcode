package problems

import (
	"fmt"
)

// Problem155 - 155. Min Stack
func Problem155() {

	minStack1 := Constructor1()

	minStack1.Push(-2)
	minStack1.Push(0)
	minStack1.Push(-3)
	fmt.Println(minStack1.GetMin())
	minStack1.Pop()
	fmt.Println(minStack1.Top())
	fmt.Println(minStack1.GetMin())

	minStack2 := Constructor1()

	minStack2.Push(2)
	minStack2.Push(1)
	fmt.Println(minStack2.Top())
	fmt.Println(minStack2.GetMin())
	minStack2.Pop()
	fmt.Println(minStack2.GetMin())
	fmt.Println(minStack2.Top())

}

// MinStack - stack structure
type MinStack struct {
	data []int
	min  int
}

// Constructor1 - Initialize your data structure here
func Constructor1() MinStack {

	stack := *new(MinStack)
	stack.data = make([]int, 0)

	return stack

}

// Push - push data
func (this *MinStack) Push(x int) {

	this.data = append(this.data, x)
	if len(this.data) == 1 || x < this.min {
		this.min = x
	}

}

// Pop - pop data
func (this *MinStack) Pop() {

	l := len(this.data)
	x := this.data[l-1]
	this.data = this.data[:l-1]
	if x == this.min && l-1 > 0 {
		this.min = this.data[0]
		for _, m := range this.data {
			if m < this.min {
				this.min = m
			}
		}
	}

}

// Top - get top data
func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

// GetMin - get min data
func (this *MinStack) GetMin() int {
	return this.min
}
