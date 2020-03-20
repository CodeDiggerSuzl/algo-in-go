package array

import (
	"fmt"
	"testing"
)

/**
Given an array, rotate the array to the right by k steps, where k is non-negative.

Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
    rotate 1 steps to the right: [7,1,2,3,4,5,6]
    rotate 2 steps to the right: [6,7,1,2,3,4,5]
    rotate 3 steps to the right: [5,6,7,1,2,3,4]

Note:
    Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
	Could you do it in-place with O(1) extra space?
*/
func rotate(nums []int, k int) {
	// make a slice
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[(i+k)%len(nums)] = nums[i]
	}
	// for i := 0; i < len(nums); i++ {
	// 	nums[i] = tmp[i]
	// }
	copy(nums, tmp)
}

var p = fmt.Print

func Test_rotate(t *testing.T) {
	ar := []int{1, 2, 3, 4, 5, 6}
	rotate(ar, 1)
}
