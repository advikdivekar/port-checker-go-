# Port Scanner Benchmark Results

Port Range: 1–1024
Runs per configuration: 1

## Target: google.com

### go_workers_100_timeout_4

#### Run 1

- Return code: 0
- Time (s): 44.19
- STDOUT:
```
Scanning google.com
Ports: 1-1024
Workers: 100
Timeout: 4.00s

[+] Port 80 is OPEN
[+] Port 443 is OPEN

========== Scan Summary ==========
Target: google.com
Ports Scanned: 1024
Workers Used: 100
Timeout: 4.00s
Open Ports: [80 443]
Timeout Count: 1022
Average Port Duration: 3993.34 ms
Total Scan Time: 44.01 seconds
==================================
```

### go_workers_500_timeout_4

#### Run 1

- Return code: 0
- Time (s): 12.19
- STDOUT:
```
Scanning google.com
Ports: 1-1024
Workers: 500
Timeout: 4.00s

[+] Port 80 is OPEN
[+] Port 443 is OPEN

========== Scan Summary ==========
Target: google.com
Ports Scanned: 1024
Workers Used: 500
Timeout: 4.00s
Open Ports: [80 443]
Timeout Count: 1022
Average Port Duration: 4000.03 ms
Total Scan Time: 12.01 seconds
==================================
```

### go_workers_1024_timeout_4

#### Run 1

- Return code: 0
- Time (s): 4.15
- STDOUT:
```
Scanning google.com
Ports: 1-1024
Workers: 1024
Timeout: 4.00s

[+] Port 80 is OPEN
[+] Port 443 is OPEN

========== Scan Summary ==========
Target: google.com
Ports Scanned: 1024
Workers Used: 1024
Timeout: 4.00s
Open Ports: [80 443]
Timeout Count: 1022
Average Port Duration: 4000.19 ms
Total Scan Time: 4.01 seconds
==================================
```

### go_workers_100_timeout_2

#### Run 1

- Return code: 0
- Time (s): 22.1
- STDOUT:
```
Scanning google.com
Ports: 1-1024
Workers: 100
Timeout: 2.00s

[+] Port 80 is OPEN
[+] Port 443 is OPEN

========== Scan Summary ==========
Target: google.com
Ports Scanned: 1024
Workers Used: 100
Timeout: 2.00s
Open Ports: [80 443]
Timeout Count: 1022
Average Port Duration: 1997.15 ms
Total Scan Time: 22.01 seconds
==================================
```

### go_workers_500_timeout_2

#### Run 1

- Return code: 0
- Time (s): 6.09
- STDOUT:
```
Scanning google.com
Ports: 1-1024
Workers: 500
Timeout: 2.00s

[+] Port 80 is OPEN

========== Scan Summary ==========
Target: google.com
Ports Scanned: 1024
Workers Used: 500
Timeout: 2.00s
Open Ports: [80]
Timeout Count: 1023
Average Port Duration: 2001.94 ms
Total Scan Time: 6.01 seconds
==================================
```

### go_workers_1024_timeout_2

#### Run 1

- Return code: 0
- Time (s): 2.12
- STDOUT:
```
Scanning google.com
Ports: 1-1024
Workers: 1024
Timeout: 2.00s


========== Scan Summary ==========
Target: google.com
Ports Scanned: 1024
Workers Used: 1024
Timeout: 2.00s
Open Ports: []
Timeout Count: 1024
Average Port Duration: 2006.47 ms
Total Scan Time: 2.01 seconds
==================================
```

## Target: scanme.nmap.org

### go_workers_100_timeout_4

#### Run 1

- Return code: 0
- Time (s): 4.09
- STDOUT:
```
Scanning scanme.nmap.org
Ports: 1-1024
Workers: 100
Timeout: 4.00s

[+] Port 22 is OPEN
[+] Port 80 is OPEN

========== Scan Summary ==========
Target: scanme.nmap.org
Ports Scanned: 1024
Workers Used: 100
Timeout: 4.00s
Open Ports: [22 80]
Timeout Count: 1
Average Port Duration: 273.60 ms
Total Scan Time: 4.00 seconds
==================================
```

### go_workers_500_timeout_4

#### Run 1

- Return code: 0
- Time (s): 4.4
- STDOUT:
```
Scanning scanme.nmap.org
Ports: 1-1024
Workers: 500
Timeout: 4.00s

[+] Port 22 is OPEN
[+] Port 80 is OPEN

========== Scan Summary ==========
Target: scanme.nmap.org
Ports Scanned: 1024
Workers Used: 500
Timeout: 4.00s
Open Ports: [22 80]
Timeout Count: 36
Average Port Duration: 1021.63 ms
Total Scan Time: 4.32 seconds
==================================
```

### go_workers_1024_timeout_4

#### Run 1

- Return code: 0
- Time (s): 4.08
- STDOUT:
```
Scanning scanme.nmap.org
Ports: 1-1024
Workers: 1024
Timeout: 4.00s

[+] Port 22 is OPEN

========== Scan Summary ==========
Target: scanme.nmap.org
Ports Scanned: 1024
Workers Used: 1024
Timeout: 4.00s
Open Ports: [22]
Timeout Count: 163
Average Port Duration: 2343.28 ms
Total Scan Time: 4.00 seconds
==================================
```

### go_workers_100_timeout_2

#### Run 1

- Return code: 0
- Time (s): 2.91
- STDOUT:
```
Scanning scanme.nmap.org
Ports: 1-1024
Workers: 100
Timeout: 2.00s

[+] Port 22 is OPEN
[+] Port 80 is OPEN

========== Scan Summary ==========
Target: scanme.nmap.org
Ports Scanned: 1024
Workers Used: 100
Timeout: 2.00s
Open Ports: [22 80]
Timeout Count: 1
Average Port Duration: 259.14 ms
Total Scan Time: 2.83 seconds
==================================
```

### go_workers_500_timeout_2

#### Run 1

- Return code: 0
- Time (s): 2.09
- STDOUT:
```
Scanning scanme.nmap.org
Ports: 1-1024
Workers: 500
Timeout: 2.00s

[+] Port 22 is OPEN
[+] Port 80 is OPEN

========== Scan Summary ==========
Target: scanme.nmap.org
Ports Scanned: 1024
Workers Used: 500
Timeout: 2.00s
Open Ports: [22 80]
Timeout Count: 135
Average Port Duration: 663.80 ms
Total Scan Time: 2.00 seconds
==================================
```

### go_workers_1024_timeout_2

#### Run 1

- Return code: 0
- Time (s): 2.09
- STDOUT:
```
Scanning scanme.nmap.org
Ports: 1-1024
Workers: 1024
Timeout: 2.00s

[+] Port 22 is OPEN
[+] Port 80 is OPEN

========== Scan Summary ==========
Target: scanme.nmap.org
Ports Scanned: 1024
Workers Used: 1024
Timeout: 2.00s
Open Ports: [22 80]
Timeout Count: 523
Average Port Duration: 1408.00 ms
Total Scan Time: 2.01 seconds
==================================
```
