package luhn

import "strings"
import "strconv"

func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	sum := 0
	if len(s) < 2 {
		return false
	}
	for j := 0; j < len(s); j++ {
		i := len(s) - j - 1
		n, err := strconv.Atoi(s[i : i+1])

		if err != nil {
			return false
		}

		if j%2 == 1 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}
	return sum%10 == 0
}
