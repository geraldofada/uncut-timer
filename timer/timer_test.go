package timer

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	timer := Start(1, "test timer")

	if timer.Id != 1 {
		t.Errorf("Expected id to be 1, got %d", timer.Id)
	}

	if timer.Name != "test timer" {
		t.Errorf("Expected name to be \"test timer\", got %s", timer.Name)
	}

	if timer.Start != time.Now() {
		t.Errorf("Expected start time to be %s, got %s", time.Now(), timer.Start)
	}

	if !timer.End.IsZero() {
		t.Errorf("Expected end time to be %s, got %s", time.Time{}, timer.End)
	}

	if timer.Elapsed != 0 {
		t.Errorf("Expected elapsed time to be 0, got %d", timer.Elapsed)
	}
}

func TestStop(t *testing.T) {
	timer := Start(1, "test timer")

	time.Sleep(5 * time.Millisecond)

	Stop(timer)

	if timer.End != time.Now() {
		t.Errorf("Expected end time to be %s, got %s", time.Now(), timer.End)
	}

	if timer.Elapsed != timer.End.Sub(timer.Start) {
		t.Errorf("Expected elapsed time to be %s, got %s",
			timer.End.Sub(timer.Start),
			timer.Elapsed)
	}
}
