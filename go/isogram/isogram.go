package isogram

import "strings"

func IsIsogram(s string) bool {
	is := true
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)
	for i := range s {
		if len(s)-1 != len(strings.Replace(s, s[i:i+1], "", -1)) {
			is = false
		}
	}
	return is
}
