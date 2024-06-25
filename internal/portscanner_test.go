// internal/portscanner_test.go
package portscanner

import (
    "testing"
    "time"
)

func TestScanPorts(t *testing.T) {
    scanner := NewScanner(1 * time.Second)
    results, err := scanner.ScanPorts("scanme.nmap.org", 22, 25)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    for _, result := range results {
        if result.Port == 22 && !result.Open {
            t.Errorf("Expected port 22 to be open")
        }
    }
}
