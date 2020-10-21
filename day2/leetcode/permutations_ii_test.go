package leetcode

import (
	"testing"
)

func helperPermUniq(nums []int, j int, path []int, ans *[][]int) {
	if j < 0 {
		copyPath := make([]int, len(path))
		copy(copyPath, path)
		*ans = append(*ans, copyPath)
	}

	used := make(map[int]bool)

	for i := j; i >= 0; i-- {
		if _, ok := used[nums[i]]; ok {
			continue
		}

		nums[j], nums[i] = nums[i], nums[j]
		//fmt.Printf("i:%v, j:%v, nums:%v\n",i, j, nums)
		path = append(path, nums[j])

		helperPermUniq(nums, j-1, path, ans)

		path = path[:len(path)-1]
		nums[j], nums[i] = nums[i], nums[j]

		used[nums[i]] = true
	}
}

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0, len(nums))
	helperPermUniq(nums, len(nums)-1, path, &ans)
	return ans
}

func TestPermuteUniq(t *testing.T) {
	//t.Log(permuteUnique([]int{1,1,2,2}))
	//t.Log(len(permuteUnique([]int{-1,2,-1,2,1,-1,2,1})))
	var s = []int{1,2,3,4,5}
	s2 := s[:3]
	s2[1] = 7
	t.Log(s)
	t.Log(s2)
}