package main

import (
	"testing"
)

func TestThousand(t *testing.T) {
	if got, want := thousand(0), "0"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(1), "1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(10), "10"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(100), "100"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(1000), "1,000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(10000), "10,000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(100000), "100,000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(1000000), "1,000,000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(10000000), "10,000,000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := thousand(100000000), "100,000,000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
