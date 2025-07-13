package core

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

type PingResult struct {
	IP     string
	Method string
	RTT    time.Duration
	TTL    int // Kann optional leer bleiben, TTL auslesen ist schwierig in pure Go
	Online bool
}

func PingSweep(subnetParts []string) []PingResult {
	var results []PingResult
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i < 255; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ipStr := fmt.Sprintf("%s.%s.%s.%d", subnetParts[0], subnetParts[1], subnetParts[2], i)

			// ICMP Ping
			pinger, err := ping.NewPinger(ipStr)
			if err == nil {
				pinger.Count = 1
				pinger.Timeout = 500 * time.Millisecond
				pinger.SetPrivileged(true)
				err = pinger.Run()
				stats := pinger.Statistics()
				if err == nil && stats.PacketsRecv > 0 {
					mu.Lock()
					results = append(results, PingResult{
						IP:     ipStr,
						Method: "ICMP",
						RTT:    stats.AvgRtt,
						Online: true,
					})
					mu.Unlock()
					fmt.Printf("Host online (ICMP): %s\n", ipStr)
					return
				}
			}

			// Fallback: TCP Port 80
			start := time.Now()
			conn, err := net.DialTimeout("tcp", ipStr+":80", 500*time.Millisecond)
			if err == nil {
				rtt := time.Since(start)
				conn.Close()
				mu.Lock()
				results = append(results, PingResult{
					IP:     ipStr,
					Method: "TCP",
					RTT:    rtt,
					Online: true,
				})
				mu.Unlock()
				fmt.Printf("Host online (TCP): %s\n", ipStr)
			}
		}(i)
	}
	wg.Wait()
	return results
}
