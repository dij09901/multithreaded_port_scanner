# PortScanner

`PortScanner` — це бібліотека на Golang для багатопоточного сканування відкритих портів.

## Використання

```go
import (
    "github.com/your-username/portscanner/internal"
)

func main() {
    scanner := portscanner.NewScanner(1 * time.Second)
    results, _ := scanner.ScanPorts("scanme.nmap.org", 80, 90)

    for _, result := range results {
        fmt.Printf("Port %d: %v\n", result.Port, result.Open)
    }
}
