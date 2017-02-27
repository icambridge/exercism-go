package pangram

import "strings"


const testVersion = 1
const abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsPangram(s string) bool {
	for count := 0; count < len(abc); count++ {
		char := string(abc[count])
		lower := strings.ToLower(char)
		if !strings.Contains(s, char) && !strings.Contains(s, lower) {
			return false
		}
	}

	return true
}
