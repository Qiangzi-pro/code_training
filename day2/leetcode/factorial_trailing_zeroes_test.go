package leetcode

import "testing"

func solutionRaw(n int) int {
	twoDivCount := 0
	fiveDivCount := 0

	for i := 1 ; i <= n; i++ {
		temp := i
		for temp % 2 == 0 {
			twoDivCount++
			temp /= 2
		}
		temp = i
		for temp % 5 == 0 {
			fiveDivCount++
			temp /= 5
		}
	}

	if twoDivCount <= fiveDivCount {
		return twoDivCount
	} else {
		return fiveDivCount
	}
}

func trailingZeroes(n int) int {
	count := 0

	for n > 0 {
		count += n/5
		n /= 5
	}

	return count
}

func TestTZ(t *testing.T) {
	t.Log(trailingZeroes(5))
}