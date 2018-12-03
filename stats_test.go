package main

import (
	"testing"
	"time"
)

func TestStats(t *testing.T) {
	s := statsHandler{}
	s.updateStats(3 * time.Second)
	s.updateStats(3 * time.Second)
	s.updateStats(3 * time.Second)
	count, average := s.getStats()
	if count != 3 {
		t.Fatalf("Expected count to be 3, got %d", count)
	}
	if average != 3000000 {
		t.Fatalf("Expected average to be 3000001, got %d", average)
	}
}
