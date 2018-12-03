package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type statsHandler struct {
	count   int64
	average int64
	lock    sync.RWMutex
}

func (s *statsHandler) updateStats(d time.Duration) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.count += 1
	s.average = s.average + (d.Nanoseconds()/1000-s.average)/s.count
}

func (s *statsHandler) getStats() (count int64, average int64) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.count, s.average
}

func (s *statsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	count, average := s.getStats()
	data := map[string]int64{"total": count, "average": average}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding stats response: %s", err)
	}
}

func statsWrapper(h http.Handler, s *statsHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(w, r)
		duration := time.Now().Sub(startTime)
		s.updateStats(duration)
	})
}
