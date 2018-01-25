// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	answer := "Whatever."
	remark = strings.Replace(remark, "\t", "", -1)
	remark = strings.Replace(remark, "\r", "", -1)
	remark = strings.Replace(remark, "\n", "", -1)
	remark = strings.Trim(remark, " ")
	isYell := strings.ToUpper(remark) == remark
	isHasChar := strings.ToUpper(remark) != strings.ToLower(remark)
	// get last character of remark
	lastChar := ""
	if len(remark) > 1 {
		lastChar = remark[len(remark)-1:]
	} else {
		lastChar = remark
	}
	if isYell && isHasChar {
		if lastChar == "?" {
			answer = "Calm down, I know what I'm doing!"
		} else {
			answer = "Whoa, chill out!"
		}
	} else {
		if lastChar == "?" {
			answer = "Sure."
		}
	}
	if len(remark) == 0 {
		answer = "Fine. Be that way!"
	}
	return answer
}
