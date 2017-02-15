// Package clause.
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	d := time.Second * 1000000000
	return t.Add(d)
}

// Reviewers don't think much of stub comments.  Replace or remove.
