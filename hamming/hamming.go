package hamming

import "fmt"

const testVersion = 5

func Distance(a, b string) (int, error) {
	aLen := len(a)
	bLen := len(b)
	if aLen != bLen {
		return -1, fmt.Errorf("both strings don't have the same length")
	}

	maxDistance := aLen

	for count := 0; count < aLen; count++ {
		if a[count] == b[count] {
			maxDistance--
		}
	}

	return maxDistance, nil
}
