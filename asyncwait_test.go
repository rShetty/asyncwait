package asyncwait

import (
	"testing"
	"time"
)

func TestAsyncWait(t *testing.T) {
	var enabled bool

	go func() {
		time.Sleep(50 * time.Millisecond)
		enabled = true
	}()

	successful := NewAsyncWait(100, 20).Check(func() bool {
		return enabled == true
	})

	if successful != true {
		t.Errorf("failed to be successful")
	}

	if enabled != true {
		t.Errorf("failed to test asyncwait, predicate was not true")
	}
}

func TestAsyncWaitFail(t *testing.T) {
	var enabled bool

	go func() {
		time.Sleep(150 * time.Millisecond)
		enabled = true
	}()

	successful := NewAsyncWait(100, 20).Check(func() bool {
		return enabled == true
	})

	if successful != false {
		t.Errorf("should have failed")
	}

	if enabled != false {
		t.Errorf("should have failed to test asyncwait, predicate should be false")
	}
}
