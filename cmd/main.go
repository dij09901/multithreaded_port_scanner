package main

import (
    "fmt"
    "log"
    ps "github.com/dij09901/portscanner/internal"
)

func main() {
    host := "scanme.nmap.org"
    startPort := 80
    endPort := 85

    scanner := ps.NewScanner(ps.DefaultTimeout)
    results, err := scanner.ScanPorts(host, startPort, endPort)
    if err != nil {
        log.Fatalf("Error scanning ports: %v", err)
    }

    for _, result := range results {
        status := "Closed"
        if result.Open {
            status = "Open"
        }
        fmt.Printf("Port %d: %s\n", result.Port, status)
    }
}
