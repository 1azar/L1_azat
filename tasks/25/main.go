package main

import (
	"fmt"
	"time"
)

// УСЛОВИЕ:
// Реализовать собственную функцию sleep.

func MySleep(duration time.Duration) {
	<-time.After(time.Duration(duration))
}

func main() {
	dur := time.Second * 5
	ts := time.Now()
	fmt.Printf("started (%s)\nwaiting %s..\n", ts, dur)
	MySleep(dur)
	fmt.Printf("stopped (%s) (%s)\n", time.Now(), time.Since(ts))
}
