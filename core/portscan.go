package core

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var DefaultPorts = []int{21, 22, 23, 25, 80, 110, 139, 143, 443, 445, 587, 8080}

func PortScanDiscovery(subnetParts []string) map[string][]int {
	results := make(map[string][]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i < 255; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ipStr := fmt.Sprintf("%s.%s.%s.%d", subnetParts[0], subnetParts[1], subnetParts[2], i)
			for _, port := range DefaultPorts {
				var address string
				if strings.Contains(ipStr, ":") {
					address = fmt.Sprintf("[%s]:%d", ipStr, port)
				} else {
					address = fmt.Sprintf("%s:%d", ipStr, port)
				}
				conn, err := net.DialTimeout("tcp", address, 400*time.Millisecond)
				if err == nil {
					conn.Close()
					mu.Lock()
					if !containsInt(results[ipStr], port) {
						results[ipStr] = append(results[ipStr], port)
					}
					mu.Unlock()
					fmt.Printf("Host online (Port %d offen): %s\n", port, ipStr)
				}
			}
		}(i)
	}
	wg.Wait()
	return results
}

// Helper für Doppelte vermeiden (Go kann das nicht “von selbst”)
func containsInt(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}
