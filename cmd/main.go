// cmd/main.go
package main

import (
    "fmt"
    "github.com/dij09901/portscanner"
    "os"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: go run main.go <host> <startPort> <endPort>")
        os.Exit(1)
    }

    host := os.Args[1]
    startPort := os.Args[2]
    endPort := os.Args[3]

    scanner := portscanner.NewScanner(portscanner.DefaultTimeout)
    results, err := scanner.ScanPorts(host, startPort, endPort)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    for _, result := range results {
        status := "Closed"
        if result.Open {
            status = "Open"
        }
        fmt.Printf("Port %d: %s\n", result.Port, status)
    }
}
