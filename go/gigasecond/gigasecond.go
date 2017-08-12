// Package gigasecond provides methods related to gigaseconds
package gigasecond

// import path for the time package from the standard library
import "time"

const (
	testVersion = 4
	gigaSecond  = 1e9 * time.Second
)

// AddGigasecond identifies a Time 1 gigasecond from the given Time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
