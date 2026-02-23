#!/bin/bash

TARGETS=("google.com" "openai.com")

WORKERS=(50 100 200 100 300 500)
TIMEOUTS=(2 2 2 4 4 4)

PORT_START=1
PORT_END=1024

mkdir -p results

RESULTS_FILE="results/benchmarks.csv"

if [ ! -f "$RESULTS_FILE" ]; then
echo "timestamp,lang,impl,workers,timeout_s,start_port,end_port,target,total_time_s" > "$RESULTS_FILE"
fi

for target in "${TARGETS[@]}"; do
  for i in {0..5}; do

    workers=${WORKERS[$i]}
    timeout=${TIMEOUTS[$i]}

    echo "Running: $target | Workers=$workers | Timeout=$timeout"

    start=$(date +%s.%N)

    go run . \
      -target $target \
      -start $PORT_START \
      -end $PORT_END \
      -workers $workers \
      -timeout $timeout

    end=$(date +%s.%N)

    duration=$(echo "$end - $start" | bc)

    echo "$(date -Iseconds),go,goroutine,$workers,$timeout,$PORT_START,$PORT_END,$target,$duration" >> "$RESULTS_FILE"

  done
done

echo "All Go benchmarks completed."