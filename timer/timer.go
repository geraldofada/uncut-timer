package timer

import "time"

// A Timer represents the main struct wich
// saves the start date and end date
type Timer struct {
	Id      int
	Name    string
	Start   time.Time
	End     time.Time
	Elapsed time.Duration
}

// Start initializes a struct Timer
// with the current date
func Start(id int, name string) *Timer {
	now := time.Now()
	t := &Timer{
		id, name, now, time.Time{}, 0,
	}

	return t
}

// Stop sets the end time and elapsed time
// of a Timer
func Stop(t *Timer) *Timer {
	now := time.Now()
	t.End = now
	t.Elapsed = t.End.Sub(t.Start)

	return t
}
