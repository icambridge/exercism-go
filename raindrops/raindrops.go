package raindrops

import "fmt"

const testVersion = 2

func Convert(i int) string {

	s := ""

	if (i % 3) == 0 {
		s = "Pling"
	}

	if (i % 5) == 0 {
		s = s + "Plang"
	}

	if (i % 7) == 0 {
		s = s + "Plong"
	}

	if s != "" {
		return s
	}

	return fmt.Sprintf("%d", i)
}
