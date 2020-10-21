package leetcode

import "testing"

func majorityElement(nums []int) int {
	size := len(nums)

	//if size < 5 {
	//	d := make(map[int]int)
	//	for _, e := range nums {
	//		d[e]++
	//	}
	//
	//	var mk, mv int
	//
	//	for key, value := range d {
	//		if value > mv {
	//			mk = key
	//			mv = value
	//		}
	//	}
	//
	//	return mk
	//}

	mid := size / 2
	j := size - 1
	i := 0
	for {
		pdx := partition(nums, i, j)
		// 奇、偶都可以满足，单独优化偶数，从时间复杂度上并不明显
		if pdx == mid {
			return nums[pdx]
		} else if pdx < mid { // 想清楚范围条件！
			i = pdx + 1
		}  else {
			j = pdx - 1
		}
	}
}

//
func partition(nums []int, i, j int) int {
	r := j - 1
	l := i
	for l <= r {
		for r >= l && nums[r] >= nums[j] {
			r--
		}

		for l <= r && nums[l] < nums[j] {
			l++
		}

		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[l], nums[j] = nums[j], nums[l]
	return l
}

func TestME(t *testing.T) {
	t.Log(majorityElement([]int{6,6,6,7,7}))
}
