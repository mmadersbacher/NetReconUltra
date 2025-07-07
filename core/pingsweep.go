package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

func PingSweep(subnetParts []string) []string {
	var online []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i < 255; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ipStr := fmt.Sprintf("%s.%s.%s.%d", subnetParts[0], subnetParts[1], subnetParts[2], i)
			pinger, err := ping.NewPinger(ipStr)
			if err != nil {
				return
			}
			pinger.Count = 1
			pinger.Timeout = 700 * time.Millisecond
			pinger.SetPrivileged(true)
			err = pinger.Run()
			if err == nil && pinger.Statistics().PacketsRecv > 0 {
				mu.Lock()
				online = append(online, ipStr)
				mu.Unlock()
				fmt.Printf("Host online (Ping): %s\n", ipStr)
			}
		}(i)
	}
	wg.Wait()
	return online
}
