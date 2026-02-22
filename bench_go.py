import subprocess
import time
from pathlib import Path

RESULTS_DIR = Path("results")
RESULTS_DIR.mkdir(exist_ok=True)

RUNS_PER_TEST = 1

TARGETS = [
    "google.com",
    "scanme.nmap.org"
    "openai.com"
]

CONFIGS = [
    {"workers": 100, "timeout": 4},
    {"workers": 500, "timeout": 4},
    {"workers": 1024, "timeout": 4},
    {"workers": 100, "timeout": 2},
    {"workers": 500, "timeout": 2},
    {"workers": 1024, "timeout": 2},
]

def run_once(target, workers, timeout):
    cmd = [
        "go", "run", ".",
        "--target", target,
        "--workers", str(workers),
        "--timeout", str(timeout),
    ]

    start = time.time()
    p = subprocess.run(cmd, text=True, capture_output=True)
    end = time.time()

    return {
        "returncode": p.returncode,
        "stdout": p.stdout.strip(),
        "stderr": p.stderr.strip(),
        "seconds": round(end - start, 2),
    }

def main():
    lines = []
    lines.append("# Port Scanner Benchmark Results\n")
    lines.append("Port Range: 1–1024")
    lines.append(f"Runs per configuration: {RUNS_PER_TEST}\n")

    for target in TARGETS:
        lines.append(f"## Target: {target}\n")

        for cfg in CONFIGS:
            name = f"go_workers_{cfg['workers']}_timeout_{cfg['timeout']}"
            lines.append(f"### {name}\n")

            for i in range(1, RUNS_PER_TEST + 1):
                result = run_once(target, cfg["workers"], cfg["timeout"])

                lines.append(f"#### Run {i}\n")
                lines.append(f"- Return code: {result['returncode']}")
                lines.append(f"- Time (s): {result['seconds']}")

                if result["stderr"]:
                    lines.append(f"- STDERR:\n```\n{result['stderr']}\n```")

                lines.append(f"- STDOUT:\n```\n{result['stdout']}\n```\n")

    out = RESULTS_DIR / "bench_go_results.md"
    out.write_text("\n".join(lines))
    print(f"Saved: {out}")

if __name__ == "__main__":
    main()