package core

import (
	"strings"
)

// Prüft, ob eine Zahl im int-Slice enthalten ist
func containsInt(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Erkennt Gerätetyp auf Basis von Banners, Hostname und Ports
func GuessDeviceType(banners map[int]string, hostname string, ports []int) string {
	for _, banner := range banners {
		s := strings.ToLower(banner)
		switch {
		case strings.Contains(s, "printer"), strings.Contains(s, "deskjet"), strings.Contains(s, "hp"):
			return "Drucker"
		case strings.Contains(s, "openwrt"), strings.Contains(s, "router"), strings.Contains(hostname, "gateway"):
			return "Router"
		case strings.Contains(s, "windows"), strings.Contains(s, "microsoft"):
			return "Windows PC"
		case strings.Contains(s, "debian"), strings.Contains(s, "ubuntu"), strings.Contains(s, "linux"):
			return "Linux/Unix"
		case strings.Contains(s, "ftp"):
			return "FTP Server"
		case strings.Contains(s, "ssh"):
			return "SSH Service"
		}
	}
	h := strings.ToLower(hostname)
	if strings.HasPrefix(h, "desktop-") {
		return "Windows PC"
	}
	if containsInt(ports, 445) && containsInt(ports, 139) {
		return "Windows/SMB Gerät"
	}
	return "Unbekannt"
}
