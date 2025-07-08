package core

import (
	"fmt"
	"log"
	"net"
	"netreconultra/models"
	"os"
	"strings"
	"time"
)

// Hilfsfunktion: Prüft, ob ein String in einem Slice enthalten ist.
func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

// Holt die Subnetz-Teile des aktiven Interfaces (z.B. ["192","168","8"])
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

// Hauptfunktion: Scan durchführen
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

	// 1. Pingsweep
	pingOnline := PingSweep(subnetParts)

	// 2. Portscan (Discovery)
	portscanResults := PortScanDiscovery(subnetParts) // map[ip][]int
	var portOnline []string
	for ip := range portscanResults {
		portOnline = append(portOnline, ip)
	}

	// 3. Hostname Discovery
	uniqueIPs := make(map[string]bool)
	for _, ip := range pingOnline {
		uniqueIPs[ip] = true
	}
	for _, ip := range portOnline {
		uniqueIPs[ip] = true
	}
	allIPs := []string{}
	for ip := range uniqueIPs {
		allIPs = append(allIPs, ip)
	}
	hostnames := HostnameDiscovery(allIPs)

	// 4. Banner Grabbing (map[string]map[string]string)
	banners := GrabBanners(portscanResults)

	// 5. Geräte-Liste erstellen
	var devices []models.Device
	for _, ip := range allIPs {
		deviceBanners := make(map[string]string)
		if b, ok := banners[ip]; ok {
			deviceBanners = b // b ist map[string]string!
		}
		dev := models.Device{
			IP:         ip,
			MAC:        "",
			Vendor:     "",
			Hostname:   hostnames[ip],
			Ports:      portscanResults[ip],
			Banners:    deviceBanners,
			FoundBy:    []string{},
			DeviceType: "",
		}
		if contains(pingOnline, ip) {
			dev.FoundBy = append(dev.FoundBy, "ping")
		}
		if contains(portOnline, ip) {
			dev.FoundBy = append(dev.FoundBy, "portscan")
		}
		dev.DeviceType = DetectDeviceType(DeviceInfo{
			IP:       ip,
			Hostname: hostnames[ip],
			Banners:  deviceBanners,
			MAC:      "",
			Ports:    portscanResults[ip],
			TTL:      0, // falls gebraucht
		})
		devices = append(devices, dev)
	}

	// 6. Ausgabe
	fmt.Println("\nScan-Ergebnis:")
	fmt.Printf("%-15s %-22s %-18s %-16s %-40s %-15s\n", "IP", "Hostname", "DeviceType", "Ports", "Banners", "FoundBy")
	for _, d := range devices {
		fmt.Printf("%-15s %-22s %-18s %-16v %-40v %-15v\n",
			d.IP, d.Hostname, d.DeviceType, d.Ports, d.Banners, d.FoundBy)
	}

	// 7. Export als JSON (immer in logs/, mit Zeitstempel + latest.json)
	now := time.Now()
	logsDir := "logs"
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		err := os.MkdirAll(logsDir, 0755)
		if err != nil {
			log.Fatalf("Konnte logs-Ordner nicht anlegen: %v", err)
		}
	}
	reportPath := fmt.Sprintf("%s/scan_%s.json", logsDir, now.Format("2006-01-02_15-04-05"))
	err := ExportToJSON(devices, reportPath)
	if err != nil {
		fmt.Println("Export-Fehler:", err)
	} else {
		fmt.Println("Report gespeichert:", reportPath)
	}
	// Immer das aktuellste JSON speichern
	_ = ExportToJSON(devices, logsDir+"/latest.json")

	// NEU: Auch ins Frontend kopieren, damit npm run dev es live anzeigen kann:
	frontendLogs := "../web/public/logs"
	if _, err := os.Stat(frontendLogs); os.IsNotExist(err) {
		_ = os.MkdirAll(frontendLogs, 0755)
	}
	_ = ExportToJSON(devices, frontendLogs+"/latest.json")
}
