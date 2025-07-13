package core

import (
	"fmt"
	"log"
	"net"
	"netreconultra/models"
	"netreconultra/utils"
	"os"
	"strings"
	"time"
)

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func extractIPs(results []PingResult) []string {
	var out []string
	for _, r := range results {
		out = append(out, r.IP)
	}
	return out
}

func getActiveSubnetParts() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		if i.Flags&net.FlagUp == 0 || i.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok || ipnet.IP.IsLoopback() || ipnet.IP.To4() == nil {
				continue
			}
			parts := strings.Split(ipnet.IP.String(), ".")
			if len(parts) < 3 {
				continue
			}
			return parts, nil
		}
	}
	return nil, fmt.Errorf("kein aktives interface mit ipv4 gefunden")
}

func RunScan() {
	fmt.Println("Starte Netzwerk-Discovery...")

	var subnetParts []string

	if len(os.Args) >= 4 {
		subnet := os.Args[3]
		subnetParts = strings.Split(subnet, ".")
		if len(subnetParts) < 3 {
			log.Fatalf("Ungültiges Subnetz (manuell): %s", subnet)
		}
		fmt.Printf("Verwende Subnetz: %s.%s.%s.0/24\n", subnetParts[0], subnetParts[1], subnetParts[2])
	} else {
		var err error
		subnetParts, err = getActiveSubnetParts()
		if err != nil {
			log.Fatalf("Fehler bei Subnetz-Erkennung: %v", err)
		}
		fmt.Printf("Verwende Subnetz: %s.%s.%s.0/24\n", subnetParts[0], subnetParts[1], subnetParts[2])
	}

	// 1. Pingsweep (ICMP)
	pingResults := PingSweep(subnetParts)
	pingOnlineIPs := extractIPs(pingResults)

	// 2. Portscan (TCP)
	portsToScan := DefaultPorts
	maxConcurrent := 100
	portMap := PortScan(pingOnlineIPs, portsToScan, maxConcurrent) // map[string][]models.PortResult

	// 3. ARP-Scan für MAC (Layer2)
	subnetCIDR := fmt.Sprintf("%s.%s.%s.0/24", subnetParts[0], subnetParts[1], subnetParts[2])
	macMap := ARPScan(subnetCIDR) // map[ip]mac

	// 4. Alle IPs sammeln (Ping, Portscan, ARP)
	uniqueIPs := make(map[string]bool)
	for _, ip := range pingOnlineIPs {
		uniqueIPs[ip] = true
	}
	for ip := range portMap {
		uniqueIPs[ip] = true
	}
	for ip := range macMap {
		uniqueIPs[ip] = true
	}
	allIPs := []string{}
	for ip := range uniqueIPs {
		allIPs = append(allIPs, ip)
	}

	// 5. Hostnames holen (Reverse Lookup)
	hostnames := HostnameDiscovery(allIPs)

	// 6. Bannergrabbing (L7)
	banners := GrabBanners(portMap)

	// 7. Devices bauen (mit MAC & Vendor)
	var devices []models.Device
	for _, ip := range allIPs {
		deviceBanners := make(map[string]string)
		if b, ok := banners[ip]; ok {
			deviceBanners = b
		}
		mac := macMap[ip]
		vendor := utils.LookupVendor(mac, "data/oui.txt")
		dev := models.Device{
			IP:         ip,
			MAC:        mac,
			Vendor:     vendor,
			Hostname:   hostnames[ip],
			Ports:      portMap[ip],
			Banners:    deviceBanners,
			FoundBy:    []string{},
			DeviceType: "",
		}
		if contains(pingOnlineIPs, ip) {
			dev.FoundBy = append(dev.FoundBy, "ping")
		}
		if len(portMap[ip]) > 0 {
			dev.FoundBy = append(dev.FoundBy, "portscan")
		}
		if mac != "" {
			dev.FoundBy = append(dev.FoundBy, "arp")
		}
		dev.DeviceType = DetectDeviceType(DeviceInfo{
			IP:       ip,
			Hostname: hostnames[ip],
			Banners:  deviceBanners,
			MAC:      mac,
			Vendor:   vendor,
			Ports:    portMap[ip],
			TTL:      0,
		})
		devices = append(devices, dev)
	}

	// 8. Elite-Ausgabe (aufgeräumt)
	fmt.Println("\nScan-Ergebnis (ELITE):")
	fmt.Printf("%-15s %-22s %-17s %-18s %-18s %-22s %-20s %-40s\n",
		"IP", "Hostname", "MAC", "Vendor", "DeviceType", "Ports", "FoundBy", "Banner")

	for _, d := range devices {
		openPorts := []string{}
		for _, p := range d.Ports {
			if p.Open {
				openPorts = append(openPorts, fmt.Sprintf("%d/%s", p.Port, p.Service))
			}
		}
		portsStr := strings.Join(openPorts, ",")
		bannerStr := ""
		for _, v := range d.Banners {
			bannerStr = v
			if len(bannerStr) > 40 {
				bannerStr = bannerStr[:40] + "..."
			}
			break
		}
		fmt.Printf("%-15s %-22s %-17s %-18s %-18s %-22s %-20s %-40s\n",
			d.IP, d.Hostname, d.MAC, d.Vendor, d.DeviceType, portsStr, strings.Join(d.FoundBy, ","), bannerStr)
	}

	// 9. Export JSON (mit scanDate & devices)
	now := time.Now()
	frontendLogs := "../../web-frontend/public/logs"
	if _, err := os.Stat(frontendLogs); os.IsNotExist(err) {
		err := os.MkdirAll(frontendLogs, 0755)
		if err != nil {
			log.Fatalf("Konnte frontend logs-Ordner nicht anlegen: %v", err)
		}
	}
	reportPath := fmt.Sprintf("%s/scan_%s.json", frontendLogs, now.Format("2006-01-02_15-04-05"))
	err := ExportToJSON(devices, reportPath)
	if err != nil {
		fmt.Println("Export-Fehler:", err)
	} else {
		fmt.Println("Report gespeichert:", reportPath)
	}
	_ = ExportToJSON(devices, frontendLogs+"/latest.json")
}
