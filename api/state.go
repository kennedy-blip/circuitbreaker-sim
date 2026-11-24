package main

import (
	"math/rand"
	"sync"
	"time"
)

type Breaker struct {
	ID                       string    `json:"id"`
	Name                     string    `json:"name"`
	Description              string    `json:"description"`
	FailureThreshold         int       `json:"failureThreshold"`
	RollingWindowSize        int       `json:"rollingWindowSize"`
	OpenCooldownSteps        int       `json:"openCooldownSteps"`
	HalfOpenMaxProbes        int       `json:"halfOpenMaxProbes"`
	HalfOpenSuccessThreshold int       `json:"halfOpenSuccessThreshold"`
	State                    string    `json:"state"`
	ConsecutiveFailures      int       `json:"consecutiveFailures"`
	OpenStepsRemaining       int       `json:"openStepsRemaining"`
	HalfOpenSuccesses        int       `json:"halfOpenSuccesses"`
	HalfOpenAttempts         int       `json:"halfOpenAttempts"`
	TotalCalls               int       `json:"totalCalls"`
	TotalFailures            int       `json:"totalFailures"`
	TotalShortCircuited      int       `json:"totalShortCircuited"`
	UpdatedAt                time.Time `json:"updatedAt"`
	window                   []bool
}
type Store struct {
	mu       sync.Mutex
	breakers map[string]*Breaker
	rng      *rand.Rand
}

var app = newStore()

func newStore() *Store {
	s := &Store{breakers: map[string]*Breaker{}, rng: rand.New(rand.NewSource(time.Now().UnixNano()))}
	now := time.Now().UTC()
	for _, b := range []*Breaker{{ID: "mpesa-stk-gateway", Name: "M-Pesa STK Gateway", Description: "Mobile money payment authorization", FailureThreshold: 5, RollingWindowSize: 10, OpenCooldownSteps: 4, HalfOpenMaxProbes: 3, HalfOpenSuccessThreshold: 2, State: "CLOSED", UpdatedAt: now}, {ID: "fx-rates-upstream", Name: "FX Rates Upstream", Description: "Foreign exchange pricing feed", FailureThreshold: 4, RollingWindowSize: 8, OpenCooldownSteps: 3, HalfOpenMaxProbes: 2, HalfOpenSuccessThreshold: 2, State: "CLOSED", UpdatedAt: now}, {ID: "search-index-cluster", Name: "Search Index Cluster", Description: "Product and catalog search", FailureThreshold: 6, RollingWindowSize: 12, OpenCooldownSteps: 5, HalfOpenMaxProbes: 3, HalfOpenSuccessThreshold: 2, State: "CLOSED", UpdatedAt: now}} {
		s.breakers[b.ID] = b
	}
	return s
}
func clone(b *Breaker) Breaker { c := *b; c.window = nil; return c }
