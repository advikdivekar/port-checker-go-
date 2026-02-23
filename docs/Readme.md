# Port Checker (Go)

A concurrent TCP connect port scanner written in Go.  
Built to study concurrency scaling, worker pool behavior, and timeout impact on large-scale port scanning.

---

## Overview

This project implements a configurable TCP port scanner that scans ports **1–1024** using a worker pool model.

It is designed to benchmark:

- Worker count scaling  
- Timeout behavior  
- Silent filtering impact  
- Total scan time convergence  

The scanner measures and reports:

- Open ports  
- Timeout count  
- Average port duration  
- Total scan time  

---

## Key Observations

- Total scan time approaches the configured timeout when most ports are silently filtered.
- Increasing worker count reduces wave stacking and improves overall completion time.
- Lower timeout values reduce total scan duration but may cause missed open ports.
- Closed ports (RST response) complete significantly faster than filtered ports.
- Concurrency overlaps I/O wait but does not reduce network latency.

---

## Project Goals

This project was created to:

- Understand TCP connect scanning behavior  
- Explore concurrency design in Go  
- Measure worker pool efficiency  
- Study timeout-driven scan convergence  
- Benchmark network behavior under controlled configurations  

---

## Scope

This implementation performs a full TCP connect scan using Go’s `net.DialTimeout`.  
It relies on the operating system’s TCP stack for retransmissions and timeout handling.

This is not a SYN (half-open) scan.

## Benchmark Results

Benchmark Results (1–1024 ports, google.com & openai.com)

• Performance scales with workers in both implementations
• Go shows slightly lower overhead at higher concurrency
• Differences are minimal for I/O-bound workloads

## Results
See: results/bench_results.md and plots/ folder

---

## Disclaimer

This project is intended strictly for educational and benchmarking purposes.  
All scans were conducted against publicly permitted test targets.

Do not scan systems without explicit authorization.