package main

import (
	"github.com/etsy/statsd"
	"time"
)

func main() {
	t1 := time.Now()

	time.Sleep(100 * time.Millisecond)

	t2 := time.Now()
	duration := int64(t2.Sub(t1) / time.Millisecond)

	STATSD := statsd.New("localhost", 8125)
	STATSD.Increment("stat.metric1")
	STATSD.Timing("createRelationship", duration)
}
