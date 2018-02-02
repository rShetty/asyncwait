package asyncwait

import "time"

// AsyncWait Represents asyncwait
type AsyncWait struct {
	tickerChan      <-chan time.Time
	maxWaitInMillis int
	doneChan        chan bool
}

// NewAsyncWait Creates a new instance of AsyncWait with maxWait and pollInterval as inputs in millis
func NewAsyncWait(maxWaitInMillis int, pollIntervalInMillis int) AsyncWait {
	return AsyncWait{
		tickerChan:      time.NewTicker(time.Millisecond * time.Duration(pollIntervalInMillis)).C,
		maxWaitInMillis: maxWaitInMillis,
		doneChan:        make(chan bool),
	}
}

// Check waits for specified maxWaitInMillis and polls every pollIntervalInMillis for the predicate truthiness
func (aw AsyncWait) Check(predicate func() bool) bool {
	go func() {
		time.Sleep(time.Millisecond * time.Duration(aw.maxWaitInMillis))
		aw.doneChan <- false
	}()

	for {
		select {
		case <-aw.tickerChan:
			go func() {
				if predicate() {
					aw.doneChan <- true
				}
			}()
		case check := <-aw.doneChan:
			return check
		}
	}
}
