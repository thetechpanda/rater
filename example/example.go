package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/thetechpanda/rater"
)

func thousand(i int64) string {
	s := strconv.Itoa(int(i))

	for i := len(s); i > 3; i -= 3 {
		s = s[:i-3] + "," + s[i-3:]
	}

	return s

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	monitor := rater.NewMonitor(ctx, time.Second)
	counters := []int64{}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case count := <-monitor.C:
				fmt.Println("events rate:", thousand(count)+"/s")
				counters = append(counters, count)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				monitor.Rate()
			}
		}
	}()

	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
	tot, count := 0, len(counters)
	for _, c := range counters {
		tot += int(c)
	}
	fmt.Println("average events rate:", thousand(int64(tot/count))+"/s")

}
