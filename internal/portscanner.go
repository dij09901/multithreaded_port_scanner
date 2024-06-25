// internal/portscanner.go
package portscanner

import (
    "fmt"
    "net"
    "sync"
    "time"
)

const (
    // DefaultTimeout defines the default timeout for port scanning
    DefaultTimeout = 1 * time.Second
)

// ScanResult represents the result of a single port scan
type ScanResult struct {
    Port    int
    Open    bool
    Message string
}

// Scanner struct holds the configuration for port scanning
type Scanner struct {
    Timeout time.Duration
}

// NewScanner returns a new instance of Scanner
func NewScanner(timeout time.Duration) *Scanner {
    return &Scanner{Timeout: timeout}
}

// ScanPorts scans the ports in the given range
func (s *Scanner) ScanPorts(host string, startPort, endPort int) ([]ScanResult, error) {
    var wg sync.WaitGroup
    results := make([]ScanResult, 0)
    resultsChan := make(chan ScanResult, endPort-startPort+1)

    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go func(p int) {
            defer wg.Done()
            result := s.scanPort(host, p)
            resultsChan <- result
        }(port)
    }

    wg.Wait()
    close(resultsChan)

    for result := range resultsChan {
        results = append(results, result)
    }

    return results, nil
}

func (s *Scanner) scanPort(host string, port int) ScanResult {
    address := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", address, s.Timeout)
    if err != nil {
        return ScanResult{Port: port, Open: false, Message: err.Error()}
    }
    conn.Close()
    return ScanResult{Port: port, Open: true, Message: "Open"}
}
