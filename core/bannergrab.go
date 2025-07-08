package core

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func BannerGrab(ip string, port int) string {
	var address string
	if strings.Contains(ip, ":") {
		address = fmt.Sprintf("[%s]:%d", ip, port)
	} else {
		address = fmt.Sprintf("%s:%d", ip, port)
	}

	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return ""
	}
	defer conn.Close()

	// Protokoll-gerechtes Verhalten (mehr Banner!)
	switch port {
	case 21, 22, 23, 25, 110, 143, 445:
		// Diese Protokolle senden oft direkt nach Verbindungsaufbau einen Banner
		// → nichts schicken, nur lesen
	case 80, 8080:
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n") // HTTP-Request für Banner
	case 587:
		// SMTP Submission, Banner kommt oft direkt
	default:
		// Bei anderen Ports: einfach probieren zu lesen
	}

	conn.SetReadDeadline(time.Now().Add(1200 * time.Millisecond))
	reader := bufio.NewReader(conn)
	var banner string
	for i := 0; i < 2; i++ { // Zwei Zeilen lesen, mehr Infos möglich
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		banner += strings.TrimSpace(line) + " "
	}
	return strings.TrimSpace(banner)
}

func GrabBanners(hosts map[string][]int) map[string]map[string]string {
	results := make(map[string]map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for ip, ports := range hosts {
		for _, port := range ports {
			wg.Add(1)
			go func(ip string, port int) {
				defer wg.Done()
				banner := BannerGrab(ip, port)
				if banner != "" {
					mu.Lock()
					if results[ip] == nil {
						results[ip] = make(map[string]string)
					}
					results[ip][fmt.Sprintf("%d", port)] = banner
					mu.Unlock()
					fmt.Printf("Banner auf %s:%d → %s\n", ip, port, banner)
				}
			}(ip, port)
		}
	}
	wg.Wait()
	return results
}
