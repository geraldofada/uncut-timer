package timer

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	timer := Start(1, "test timer")

	if timer.id != 1 {
		t.Errorf("Expected id to be 1, got %d", timer.id)
	}

	if timer.name != "test timer" {
		t.Errorf("Expected name to be \"test timer\", got %s", timer.name)
	}

	if timer.start != time.Now() {
		t.Errorf("Expected start time to be %s, got %s", time.Now(), timer.start)
	}

	if !timer.end.IsZero() {
		t.Errorf("Expected end time to be %s, got %s", time.Time{}, timer.end)
	}

	if timer.elapsed != 0 {
		t.Errorf("Expected elapsed time to be 0, got %d", timer.elapsed)
	}
}

func TestStop(t *testing.T) {
	timer := Start(1, "test timer")

	time.Sleep(5 * time.Millisecond)

	Stop(timer)

	if timer.end != time.Now() {
		t.Errorf("Expected end time to be %s, got %s", time.Now(), timer.end)
	}

	if timer.elapsed != timer.end.Sub(timer.start) {
		t.Errorf("Expected elapsed time to be %s, got %s",
			timer.end.Sub(timer.start),
			timer.elapsed)
	}
}
