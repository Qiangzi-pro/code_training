package leetcode

import (
	"strconv"
	"testing"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	operators := map[string]func(e1, e2 int) int{
		"+": func(e1, e2 int) int {
			return e1 + e2
		},
		"-": func(e1, e2 int) int {
			return e1 - e2
		},
		"*": func(e1, e2 int) int {
			return e1 * e2
		},
		"/": func(e1, e2 int) int {
			return e1 / e2
		},
	}

	for _, v := range tokens {
		var elem int
		if operFunc, ok := operators[v]; ok {
			size := len(stack)
			e1, e2 := stack[size-2], stack[size-1]
			stack = stack[:size-2]
			elem = operFunc(e1, e2)
		} else {
			elem, _ = strconv.Atoi(v)
		}
		stack = append(stack, elem)
	}

	return stack[0]
}

func TestIntegerDivide(t *testing.T) {
	t.Log(13 / 5)
}
