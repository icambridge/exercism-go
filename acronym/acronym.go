package acronym

import (
	"strings"
	"fmt"
)

const testVersion = 2

func Abbreviate(s string) string {
	var letters []string
	parts := strings.Split(s, " ")
	for _, part := range parts {
		subParts := strings.Split(part, "-")
		for _, subPart := range subParts {
			l := strings.ToUpper(string(subPart[0]))
			letters = append(letters, l)
			partLen := len(subPart)

			var subletters []string
			found := 0
			howMany := partLen-2
			for count := 1; count < partLen-1; count++ {
				o := fmt.Sprintf("%s", string(subPart[count]))
				u := strings.ToUpper(o)
				if u == o {
					found++
					subletters = append(subletters, o)
				}
			}
			if howMany != found {
				letters = append(letters, subletters...)
			}
		}
	}

	return strings.Join(letters, "")
}
