package leetcode

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	return twoPointerSolution(nums)
}

func binarySearch(val int, nums []int, i int) int {
	j := len(nums) - 1

	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] < val {
			i = mid + 1
		} else if nums[mid] > val {
			j = mid - 1
		} else {
			return mid
		}
	}

	return -1
}


// two pointer solution
func twoPointerSolution(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		for i != 0 && i < len(nums) && nums[i]==nums[i-1] {
			i++
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func TestSum(t *testing.T) {
	t.Log(threeSum([]int{-1,0,1,2,-1,-4}))
	//var d [][]int
	//t.Log(len(d))
}