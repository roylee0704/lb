package lb

import (
	"fmt"
	"math/rand"
	"time"
)

// request is a job request to be sent off to load-balancer.
type request struct {
	fn func() int // work to be performed by worker
	c  chan int   // report back when the work's completed.
}

func workFn() int {
	fmt.Println("working...")
	return rand.Intn(10)
}

var (
	waitSleep = time.Sleep
	waitDone  = func() {} // noop
)

// Requester simulately sleeps for  2 * nWorkers seconds, and sends request
// via work channel.
func Requester(work chan request) {
	c := make(chan int)
	for {
		waitSleep(time.Duration(rand.Int63n(2 * 2e9))) // spend time
		work <- request{workFn, c}                     // send request out
		result := <-c                                  // waits for result
		fmt.Println(result)                            // process result, whatever.
		waitDone()
	}
}
