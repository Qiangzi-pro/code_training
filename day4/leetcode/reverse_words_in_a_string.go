package leetcode

import (
	"strings"
)

func reverseWords(s string) string {
	//'dlrow olleh'
	var builder strings.Builder
	const Space = ' '

	for i := 0; i < len(s); {
		if s[i] == Space {
			i++
			continue
		}

		var j = i + 1
		for ; j < len(s); j++ {
			if s[j] == Space {
				break
			}
		}

		if builder.Len() != 0 {
			builder.WriteRune(Space)
		}

		builder.WriteString(s[i : j])
		i = j
	}

	s2 := builder.String()
	return s2
}
