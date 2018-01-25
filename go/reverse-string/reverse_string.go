package reverse

import "fmt"

func String(s string) string {
	if s != "" {
		return String(s[1:]) + s[:1]
	} else {
		return ""
	}
}
