package twofer

import "strings"

func ShareWith(name string) string {
	msg := "One for X, one for me."
	switch strings.ToLower(name) {
	case "you", "":
		msg = strings.Replace(msg, "X", "you", -1)
		break
	default:
		msg = strings.Replace(msg, "X", name, -1)
		break
	}
	return msg
}
