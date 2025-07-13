package core

import (
	"fmt"
	"net"
	"netreconultra/models"
	"sync"
	"time"
)

var DefaultPorts = []int{21, 22, 23, 25, 80, 110, 139, 143, 443, 445, 587, 8080}

// guessService: Gibt Dienstnamen anhand Portnummer zurück
func guessService(port int) string {
	services := map[int]string{
		21: "FTP", 22: "SSH", 23: "Telnet", 25: "SMTP",
		80: "HTTP", 110: "POP3", 139: "NetBIOS", 143: "IMAP",
		443: "HTTPS", 445: "SMB", 587: "SMTP-SSL", 8080: "HTTP-Alt",
	}
	if name, ok := services[port]; ok {
		return name
	}
	return "Unknown"
}

// PortScan: Gibt dir ein map[string][]models.PortResult – das ist Elite-Level
func PortScan(ipList []string, ports []int, maxConcurrent int) map[string][]models.PortResult {
	results := make(map[string][]models.PortResult)
	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, maxConcurrent)

	for _, ip := range ipList {
		ip := ip // für goroutine capture!
		for _, port := range ports {
			port := port
			wg.Add(1)
			go func() {
				defer wg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()

				addr := fmt.Sprintf("%s:%d", ip, port)
				start := time.Now()
				conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
				duration := time.Since(start)

				open := false
				if err == nil {
					conn.Close()
					open = true
				}

				portResult := models.PortResult{
					Port:     port,
					Open:     open,
					Service:  guessService(port),
					Duration: duration,
					Banner:   "",
				}

				mu.Lock()
				results[ip] = append(results[ip], portResult)
				mu.Unlock()

				if open {
					fmt.Printf("Host %s → Port %d offen (%s)\n", ip, port, guessService(port))
				}
			}()
		}
	}
	wg.Wait()
	return results
}
