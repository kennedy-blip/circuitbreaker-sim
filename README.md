# CircuitBreakerSim

A Go-powered circuit-breaker state-machine and failover simulator with a zero-build Vanilla JavaScript frontend.

## Features
- CLOSED / OPEN / HALF_OPEN state machine, failure thresholds and cooldowns
- Concurrent simulated traffic, configurable error rates and upstream latency
- Fast-fail, optional fallback responses and half-open recovery probes
- Pre-seeded M-Pesa STK, FX rates and search-index dependency profiles
- Tune/reset breakers and add custom profiles
- Offline client-side fallback simulator

## Local checks
Run gofmt -w api/*.go, go vet ./api/... and go test ./... from the repository root. Use Vercel CLI with vercel dev for local serverless development.

## Deploy
Import kennedy-blip/circuitbreaker-sim into Vercel. Static assets live in public and Go handlers live in api. No Node build step is used.

## Runtime limitation
The in-memory registry is for simulation and is not durable across serverless instances or cold starts. Use shared persistent storage for real production workloads. The simulator never contacts real payment or financial services.
