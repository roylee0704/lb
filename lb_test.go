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
	<-sleep

	req := <-work
	result := req.fn()

	req.c <- result

	<-done

}
