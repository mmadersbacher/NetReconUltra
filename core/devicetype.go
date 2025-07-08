package core

import (
	"strings"
)

type DeviceInfo struct {
	IP       string
	Hostname string
	Banners  map[string]string
	MAC      string
	Ports    []int
	TTL      int
}

func PortOpen(info DeviceInfo, port int) bool {
	for _, p := range info.Ports {
		if p == port {
			return true
		}
	}
	return false
}

func DetectDeviceType(info DeviceInfo) string {
	h := strings.ToLower(info.Hostname)
	banner := ""
	for _, b := range info.Banners {
		banner += strings.ToLower(b) + " "
	}

	// Router
	if strings.Contains(h, "gateway") || strings.Contains(h, "router") {
		return "Router"
	}
	// Drucker
	if strings.Contains(h, "hp") || strings.Contains(h, "printer") || strings.Contains(banner, "hp http server") {
		return "Drucker"
	}
	// Windows PC/Server (SMB)
	if PortOpen(info, 445) && PortOpen(info, 139) && info.TTL >= 120 {
		return "Windows PC"
	}
	// NAS
	if PortOpen(info, 2049) || PortOpen(info, 5000) || strings.Contains(h, "nas") {
		return "NAS"
	}
	// Linux/Unix
	if (info.TTL >= 60 && info.TTL < 70) || strings.Contains(banner, "linux") {
		return "Linux/Unix"
	}
	return "Unbekannt"
}
