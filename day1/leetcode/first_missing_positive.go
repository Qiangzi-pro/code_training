package leetcode

import (
	"sort"
)

// O(nlogn)的解法
func solutionRaw(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 1
	}

	sort.Ints(nums)

	idx := len(nums)-1

	for idx >= 0 && nums[idx] > 0 {
		idx--
	}

	if idx == len(nums) - 1 {
		return 1
	}

	pInt := 1
	var j int
	for j = idx+1; j < len(nums); {
		for j+1 < len(nums) && nums[j] == nums[j+1] {
			j++
		}
		if nums[j] == pInt {
			j++
			pInt++
		} else {
			return pInt
		}
	}

	return pInt
}

// hash on site
func hashOnSite(nums []int) int {

	for idx := 0; idx < len(nums); {
		v := nums[idx]
		if v-1 == idx || v < 1 || v > len(nums)+1 || nums[v-1] == nums[idx]{
			idx++
			continue
		}
		nums[v-1], nums[idx] = nums[idx], nums[v-1]
	}

	for idx, v := range nums {
		if v-1 != idx {
			return idx + 1
		}
	}
	return len(nums)+1
}

func firstMissingPositive(nums []int) int {
	return hashOnSite(nums)
}