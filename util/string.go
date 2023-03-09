package util

import (
	"strings"
	"unicode"
)

// UpperFirstLetter Capitalization of all initials, if keep true keep separator
func UpperFirstLetter(s, separator string, keep bool) string {
	var result string
	if s != "" {
		temp := strings.Split(s, separator)
		for y := 0; y < len(temp); y++ {
			vv := []rune(temp[y])
			for i := 0; i < len(vv); i++ {
				// 不是大写才转换
				if i == 0 && !unicode.IsUpper(vv[i]) {
					vv[i] -= 32
				}
				result += string(vv[i]) // + string(vv[i+1])
			}
			// 保持分隔符
			if keep {
				result += separator
			}
		}
	}

	return result
}
