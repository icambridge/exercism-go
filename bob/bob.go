package bob // package name must match the package name in bob_test.go

import (
	"strings"
)

const testVersion = 2 // same as targetTestVersion


func Hey(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "Fine. Be that way!"
	}

	var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
						  "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		                  "ä", "ö", "ü"}

	numberOfLetters := 0
	numberOfCaps    := 0
	for _, c := range strings.Split(s, "") {
		for _, l := range arr {
			if strings.ToLower(c) == l {
				numberOfLetters++
				if c == strings.ToUpper(c) {
					numberOfCaps++
				}
			}
		}
	}

	if numberOfLetters != 0 && numberOfCaps == numberOfLetters {
		return "Whoa, chill out!"
	}
	p := strings.LastIndex(s, "?")
	length := len(s) - 1
	if p == length {
		return "Sure."
	}

	return "Whatever."
}