package leetcode

import "testing"

func IndexOf(list []rune, v rune) int {
	for i, c := range list {
		if c == v {
			return i
		}
	}
	return -1
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	bracketMap := map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}

	for _, c := range s {
		if v, ok := bracketMap[c]; ok {
			size := len(stack)
			if size > 0 {
				if stack[size-1] == v {
					stack = stack[:size-1]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func TestValid(t *testing.T) {
	s := "evolution"
	t.Log(len(s))
}


