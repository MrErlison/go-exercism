package bob

import (
	"regexp"
	"strings"
)

// Hey returns what Bob will reply to someone
// when they say something to him or ask him a question.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	isUpper := strings.ToUpper(remark) == remark
	isQuestion := strings.HasSuffix(remark, "?")
	hasLetter, err := regexp.MatchString(`[[:alpha:]]`, remark)
	if err != nil {
		return ""
	}

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isQuestion:
		if isUpper && hasLetter {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	case isUpper && hasLetter:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
