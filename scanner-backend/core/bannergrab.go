// Datei: core/bannergrab.go
package core

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"netreconultra/models"
	"strings"
	"sync"
	"time"
)

// BannerGrab: Liest Layer 7 Banner eines Ports, inklusive HTTP Server Header
func BannerGrab(ip string, port int, tlsEnabled bool) string {
	address := fmt.Sprintf("%s:%d", ip, port)
	timeout := 1200 * time.Millisecond

	var conn net.Conn
	var err error

	if tlsEnabled {
		conn, err = tls.DialWithDialer(&net.Dialer{Timeout: timeout}, "tcp", address, &tls.Config{InsecureSkipVerify: true})
	} else {
		conn, err = net.DialTimeout("tcp", address, timeout)
	}
	if err != nil {
		return ""
	}
	defer conn.Close()

	switch port {
	case 21, 22, 23, 25, 110, 143, 445:
		// Banner kommt nach Verbindung
	case 80, 8080, 443:
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	default:
		// Sonstige Ports einfach lesen
	}

	conn.SetReadDeadline(time.Now().Add(timeout))
	reader := bufio.NewReader(conn)

	var bannerLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		bannerLines = append(bannerLines, strings.TrimSpace(line))
	}

	// Suche Server: Header
	for _, line := range bannerLines {
		if strings.HasPrefix(strings.ToLower(line), "server:") {
			return line // z.B. "Server: nginx/1.18.0"
		}
	}

	// Falls kein Server-Header, gib kompletten Banner (1. Zeile) zurück
	if len(bannerLines) > 0 {
		return bannerLines[0]
	}

	return ""
}

// GrabBanners: Macht paralleles Bannergrabbing für alle IPs & Ports
func GrabBanners(portMap map[string][]models.PortResult) map[string]map[string]string {
	results := make(map[string]map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for ip, ports := range portMap {
		for _, pr := range ports {
			wg.Add(1)
			go func(ip string, pr models.PortResult) {
				defer wg.Done()
				tlsEnabled := false
				if pr.Port == 443 {
					tlsEnabled = true
				}
				banner := BannerGrab(ip, pr.Port, tlsEnabled)
				if banner != "" {
					mu.Lock()
					if results[ip] == nil {
						results[ip] = make(map[string]string)
					}
					results[ip][fmt.Sprintf("%d", pr.Port)] = banner
					mu.Unlock()
					fmt.Printf("Banner auf %s:%d → %s\n", ip, pr.Port, banner)
				}
			}(ip, pr)
		}
	}
	wg.Wait()
	return results
}
