package timer

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	timer := Start("test timer")

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

	if timer.Stopped != false {
		t.Errorf("Expected Stopped filed to be false, got %t", timer.Stopped)
	}
}

func TestStop(t *testing.T) {
	timer := Start("test timer")

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

	if timer.Stopped != true {
		t.Errorf("Expected stopped field to be true, got %t", timer.Stopped)
	}
}
