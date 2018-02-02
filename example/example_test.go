package example

import (
	"testing"

	"github.com/rshetty/asyncwait"
)

func TestDoWorkAsyncFunction(t *testing.T) {
	inChan := make(chan int, 5)
	outChan := make(chan int, 5)
	maxWaitInMillis := 100
	pollIntervalInMillis := 20

	doWork(inChan, outChan)

	successful := asyncwait.NewAsyncWait(maxWaitInMillis, pollIntervalInMillis).Check(func() bool {
		return len(inChan) == 5
	})

	if successful != true {
		t.Errorf("failed to assert async function")
	}

	if <-outChan != 3 {
		t.Errorf("failed to assert out channel")
	}
}
