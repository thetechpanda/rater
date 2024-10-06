package rater_test

import (
	"context"
	"testing"
	"time"

	"github.com/thetechpanda/rater"
)

func TestMonitor(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	monitor := rater.NewMonitor(ctx, time.Second/3)
	for i := 0; i < 10; i++ {
		monitor.Rate()
	}
	if got := monitor.Value(); got != 10 {
		t.Fatalf("got %d, want %d", got, 10)
	}
	time.Sleep(time.Second/3 + 10*time.Millisecond)
	if monitor.Value() != 0 {
		t.Fatalf("got %d, want %d", monitor.Value(), 0)
	}
}

func TestMonitorChannel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	monitor := rater.NewMonitor(ctx, time.Second/3)
	for i := 0; i < 10; i++ {
		monitor.Rate()
	}
	value := <-monitor.C
	if value != int64(10) {
		t.Fatalf("got %d, want %d", value, 10)
	}
	if got := monitor.Value(); got != 0 {
		t.Fatalf("got %d, want %d", got, 0)
	}
}
