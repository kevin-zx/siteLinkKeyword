package han

import (
	"regexp"
	"unicode"
)

// IsHan can judge a string, whether it is a legal Chinese word and english word.
func IsHan(str string) bool  {
	for _, r := range str {
		if !(unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r)))||isNumOrDigit(string(r))) {
			return false
		}
	}
	return true
}

var nums = regexp.MustCompile("[0-9]")
var letters =  regexp.MustCompile("[a-zA-Z]")
func isNumOrDigit(str string) bool {
	if nums.MatchString(str) {
		return true
	}
	if letters.MatchString(str) {
		return true
	}
	return false
}
