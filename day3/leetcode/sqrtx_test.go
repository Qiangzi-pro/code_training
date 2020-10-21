package leetcode

import (
	"math"
	"testing"
)

func precisionBinary(x float64) int {
	var i, j = 0.0, x
	precision := 0.01

	for {
		mid := i + (j-i)/2
		square := mid * mid
		diff := square - x
		if math.Abs(diff) <= precision {
			return int(math.Floor(mid))
		} else if diff < -precision {
			i = mid
		} else if diff > precision {
			j = mid
		}
	}
}

// a^2 <= x < (a+1)^2
func squareBinary(x int) int {
	var i, j = 0, x

	for {
		mid := i + (j-i)/2
		square := mid * mid
		nextSqu := (mid+1) * (mid+1)
		if square <= x && nextSqu > x {
			return mid
		} else if nextSqu <= x{
			i = mid
		} else if square > x {
			j = mid
		}
	}
}

func mySqrt(x int) int {
	// better
	var i, j, ans = 0, x, -1

	for i <= j {
		mid := i + (j-i)/2
		if mid*mid <= x {
			ans = mid
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return ans
}

func TestSqrt(t *testing.T) {
	t.Log(mySqrt(10))
}