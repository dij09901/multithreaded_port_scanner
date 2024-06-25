# PortScanner

`PortScanner` — is a Golang library for multi-threaded open port scanning.

## Application of the code

```go
import (
    "github.com/dij09901/portscanner/internal"
)

func main() {
    scanner := portscanner.NewScanner(1 * time.Second)
    results, _ := scanner.ScanPorts("scanme.nmap.org", 80, 90)

    for _, result := range results {
        fmt.Printf("Port %d: %v\n", result.Port, result.Open)
    }
}

```


```
portscanner/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── portscanner.go
│   └── portscanner_test.go
│
├── go.mod
└── README.md
```