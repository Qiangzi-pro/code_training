package leetcode

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func longestValidParentheses(s string) int {
	d := make([]int, len(s))
	stack := make([]int, 0, 100)

	for idx, c := range s {
		if c == '(' {
			stack = append(stack, idx)
		} else {
			size := len(stack)
			if size > 0 {
				leftBracket := stack[size-1]
				stack = stack[:size-1]

				leftNei, rightNei := leftBracket-1, leftBracket+1
				var cand1, cand2 int
				if leftNei > -1 {
					cand1 = d[leftNei]
				}

				if rightNei < len(s) {
					cand2 =  d[rightNei]
				}

				count := cand1 + cand2 + 1
				d[leftBracket] = count
				d[idx] = count

				for leftNei > -1 && d[leftNei] > 0 {
					d[leftNei] = count
					leftNei--
				}
			}
		}
	}

	count := 0

	for _, v := range d {
		if v > count {
			count = v
		}
	}

	return count * 2
}
