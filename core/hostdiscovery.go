package core

import (
	"fmt"
	"net"
)

func HostnameDiscovery(ips []string) map[string]string {
	results := make(map[string]string)
	for _, ip := range ips {
		names, err := net.LookupAddr(ip)
		if err == nil && len(names) > 0 {
			hostname := names[0]
			fmt.Printf("Host %s → Hostname: %s\n", ip, hostname)
			results[ip] = hostname
		} else {
			fmt.Printf("Host %s → Kein Hostname (rDNS) gefunden\n", ip)
			results[ip] = ""
		}
	}
	return results
}
