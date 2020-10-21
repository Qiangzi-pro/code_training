package leetcode

import "testing"

func majorityElement2(nums []int) []int {
	//
	var cand1, count1 int
	var cand2, count2 int

	for _, v := range nums {
		if count1 == 0 {
			if count2 == 0 || (count2 != 0 && cand2 != v) {
				count1 = 1
				cand1 = v
			} else if cand2 == v { // cand2 == v
				count2++
			}
		} else if cand1 == v {
			count1++
		} else { // count1 != 0
			if count2 == 0 {
				count2 = 1
				cand2 = v
			} else if cand2 == v {
				count2++
			} else { // count2 != 0
				count1--
				count2--
			}
		}
	}

	ret := make([]int, 0, 2)

	if count1 > 0 {
		stat1 := statElemCount(cand1, nums)
		if stat1 > len(nums) / 3 {
			ret = append(ret, cand1)
		}
	}

	if count2 > 0 {
		stat2 := statElemCount(cand2, nums)
		if stat2 > len(nums) / 3 {
			ret = append(ret, cand2)
		}
	}

	return ret
}

func statElemCount(ele int, nums []int) int {
	stat := 0
	for _, v := range nums {
		if v == ele {
			stat++
		}
	}
	return stat
}

func TestME2(t *testing.T) {
	t.Log(majorityElement2([]int{3,2,3}))
}
