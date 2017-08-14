//Package clock provides methods for doing clocky things
package clock

import "fmt"

const testVersion = 4
const day = 24 * 60

type Clock int

// New returns a Clock given an hour and minute
func New(hour, minute int) Clock {
	t := (hour*60 + minute) % day

	if t < 0 {
		t += day
	}

	return Clock(t)
}

// String expresses the given Clock as a string
func (c Clock) String() string {
	h := c / 60
	m := c % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

// Add adds the given number of minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}
