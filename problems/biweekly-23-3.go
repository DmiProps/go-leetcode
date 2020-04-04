package problems

import "fmt"

// Biweekly23_3 - test 3
func Biweekly23_3() {

	fmt.Println(checkOverlap(1, 0, 0, 1, -1, 3, 1))
	fmt.Println(checkOverlap(1, 0, 0, -1, 0, 0, 1))
	fmt.Println(checkOverlap(1, 1, 1, -3, -3, 3, 3))
	fmt.Println(checkOverlap(1, 1, 1, 1, -3, 2, -1))
	fmt.Println(checkOverlap(4, 102, 50, 0, 0, 100, 100))

	fmt.Println(checkOverlap(2, 8, 6, 5, 1, 10, 4))

}

func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {

	x1 -= x_center
	x2 -= x_center
	y1 -= y_center
	y2 -= y_center

	if x1 <= 0 && y1 <= 0 && x2 >= 0 && y2 >= 0 {
		return true
	}
	if y1 <= 0 && y2 >= 0 {
		if x1 <= 0 {
			return -x2 <= radius
		}
		return x1 <= radius
	}
	if x1 <= 0 && x2 >= 0 {
		if y1 <= 0 {
			return -y2 <= radius
		}
		return y1 <= radius
	}
	min := (x1*x1 + y1*y1)
	m := (x2*x2 + y2*y2)
	if m < min {
		min = m
	}
	m = (x1*x1 + y2*y2)
	if m < min {
		min = m
	}
	m = (x2*x2 + y1*y1)
	if m < min {
		min = m
	}
	return min <= (radius * radius)

}
