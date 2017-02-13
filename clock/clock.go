// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import (
	"fmt"
	"time"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {

	t := getTime(hour, minute)

	return Clock{t.Hour(), t.Minute()}
}

func (c Clock) String() string {

	t := getTime(c.hour, c.minute)
	return t.Format("15:04")
}

func (c Clock) Add(minutes int) Clock {
	s := fmt.Sprintf("%dm", minutes)
	d, err := time.ParseDuration(s)

	if err != nil {
		return c
	}
	t := getTime(c.hour, c.minute)
	newTime := t.Add(d)

	return Clock{newTime.Hour(), newTime.Minute()}
}

func getTime(hour, minute int) time.Time {

	return time.Date(2014, time.December, 31, hour, minute, 0, 0, time.UTC)
}
