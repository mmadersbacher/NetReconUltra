package core

import (
	"fmt"
	"log"
	"net"
	"net/netip"
	"sync"
	"time"

	"github.com/mdlayher/arp"
)

// ARPScan: Find all MACs in Subnetz, returns map[ip]mac – PARALLEL & mit Timeout!
func ARPScan(subnet string) map[string]string {
	results := make(map[string]string)
	ifi, err := getPrimaryInterface()
	if err != nil {
		log.Println("Kein Interface gefunden für ARP:", err)
		return results
	}

	ip, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		log.Println("Subnetz-Parsing Fehler:", err)
		return results
	}

	// Da RAW-Socket shared ist, pro Thread neue Connection!
	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, 24) // max 24 parallele ARP-Requests

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ipCopy := net.IP(make([]byte, len(ip)))
		copy(ipCopy, ip)
		wg.Add(1)
		go func(ip net.IP) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			conn, err := arp.Dial(ifi)
			if err != nil {
				return
			}
			defer conn.Close()
			ipAddr, ok := netip.AddrFromSlice(ip)
			if !ok {
				return
			}
			macCh := make(chan string, 1)
			go func() {
				mac, err := conn.Resolve(ipAddr)
				if err == nil {
					macCh <- mac.String()
				} else {
					macCh <- ""
				}
			}()
			select {
			case mac := <-macCh:
				if mac != "" {
					mu.Lock()
					results[ip.String()] = mac
					mu.Unlock()
				}
			case <-time.After(350 * time.Millisecond):
				// Timeout, keine MAC gefunden
			}
		}(ipCopy)
	}
	wg.Wait()
	return results
}

// increment IP (byteweise)
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// Finde primäres Netzwerkinterface (mit IP und up)
func getPrimaryInterface() (*net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
				return &iface, nil
			}
		}
	}
	return nil, fmt.Errorf("kein Interface gefunden")
}
