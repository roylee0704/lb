package lb

import (
	"testing"
	"time"
)

func TestRequester(t *testing.T) {
	work := make(chan request)
	sleep := make(chan struct{})
	done := make(chan struct{})

	waitSleep = func(d time.Duration) {
		sleep <- struct{}{}
	}
	waitDone = func() {
		done <- struct{}{}
	}

	go Requester(work)
	<-sleep            // to release from sleep
	req := <-work      // to obtain request sent by requester
	result := req.fn() // to processs job in request
	req.c <- result    // and report job result back to requester
	<-done             // wait until result is printed.
}
