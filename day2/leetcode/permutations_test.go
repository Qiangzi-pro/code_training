package leetcode

import (
	"fmt"
	"testing"
)

// 这里可以加深对 []int 是一个引用类型的理解
// 深度优先遍历，& 状态重置 -> 回溯
func helper(nums []int, i int, path []int, ans *[][]int) {
	if i > len(nums)-1 {
		copyOrder := make([]int, len(path))
		copy(copyOrder, path)
		*ans = append(*ans, copyOrder)
		return
	}

	for k := i; k < len(nums); k++ {
		nums[i], nums[k] = nums[k], nums[i]

		path = append(path, nums[i])
		helper(nums, i+1, path, ans)

		path = path[:len(path)-1]
		nums[i], nums[k] = nums[k], nums[i]
	}
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	order := make([]int, 0, len(nums))
	helper(nums, 0, order, &ans)
	fmt.Println(order)
	return ans
}

func TestPermute(t *testing.T) {
	t.Log(permute([]int{1,2,3}))
}