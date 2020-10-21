package leetcode

import "testing"

func findKthLargest(nums []int, k int) int {
	var i, j = 0, len(nums) - 1
	kthIndex := len(nums) - k
	for {
		pth := partition(nums, i, j)
		if pth == kthIndex {
			return nums[pth]
		} else if pth < kthIndex {
			i = pth + 1
		} else {
			j = pth - 1
		}
	}
}

func partition(nums []int, i, j int) int {
	l, r := i, j-1

	for l <= r {
		for l <= r && nums[r] > nums[j] {
			r--
		}

		for l <= r && nums[l] <= nums[j] {
			l++
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			r--
			l++
		}
	}

	nums[l], nums[j] = nums[j], nums[l]
	return l
}

func TestM(t *testing.T) {
	t.Log(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
}