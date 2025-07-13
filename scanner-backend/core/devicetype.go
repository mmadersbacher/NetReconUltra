package core

import (
	"netreconultra/models"
	"regexp"
	"strings"
)

// PortOpen prüft, ob ein bestimmter Port offen ist.
func PortOpen(ports []models.PortResult, port int) bool {
	for _, p := range ports {
		if p.Port == port && p.Open {
			return true
		}
	}
	return false
}

// DeviceInfo fasst alle relevanten Daten für die Device-Erkennung zusammen.
type DeviceInfo struct {
	IP       string
	Hostname string
	Banners  map[string]string
	MAC      string
	Vendor   string
	Ports    []models.PortResult
	TTL      int
}

// DetectDeviceType wertet alle Infos mit einem gewichteten Punktesystem aus und gibt den wahrscheinlichsten Typ zurück.
func DetectDeviceType(info DeviceInfo) string {
	scores := map[string]int{}

	addPoints := func(deviceType string, pts int) {
		scores[deviceType] += pts
	}

	h := strings.ToLower(info.Hostname)
	vendor := strings.ToLower(info.Vendor)

	// Kombiniere alle Banner in einen String (klein)
	bannerAll := ""
	for _, b := range info.Banners {
		bannerAll += strings.ToLower(b) + " "
	}

	// Vendor (hoch gewichtet)
	if strings.Contains(vendor, "vmware") || strings.Contains(vendor, "virtualbox") {
		addPoints("Virtuelle Maschine", 15)
	}
	if strings.Contains(vendor, "cisco") {
		addPoints("Cisco Router/Device", 14)
	}
	if strings.Contains(vendor, "hewlett") || strings.Contains(vendor, "hp") {
		addPoints("HP Drucker/Device", 14)
	}
	if strings.Contains(vendor, "raspberry") {
		addPoints("Raspberry Pi", 14)
	}
	if strings.Contains(vendor, "apple") {
		addPoints("Apple Gerät", 13)
	}
	if strings.Contains(vendor, "samsung") {
		addPoints("Samsung Smartphone/Device", 12)
	}

	// TTL (mittel bis hoch)
	switch info.TTL {
	case 128:
		addPoints("Windows Gerät", 11)
		addPoints("Windows/SMB Gerät", 8)
	case 64:
		addPoints("Linux/Unix Gerät", 11)
	case 255:
		addPoints("Cisco/Apple Gerät", 10)
	}

	// Ports (mittel)
	if PortOpen(info.Ports, 445) || PortOpen(info.Ports, 139) {
		addPoints("Windows/SMB Gerät", 12)
	}
	if (PortOpen(info.Ports, 22) && (PortOpen(info.Ports, 80) || PortOpen(info.Ports, 8080))) || PortOpen(info.Ports, 443) {
		addPoints("Linux/Unix Webserver", 11)
	}
	if PortOpen(info.Ports, 9100) || PortOpen(info.Ports, 515) {
		addPoints("Drucker/Print Server", 11)
	}
	if PortOpen(info.Ports, 2049) || PortOpen(info.Ports, 5000) || strings.Contains(h, "nas") {
		addPoints("NAS", 12)
	}

	// Banner (mittel bis hoch), jetzt mit Webserver Erkennung
	for _, b := range info.Banners {
		bLower := strings.ToLower(b)
		switch {
		case strings.Contains(bLower, "apache"):
			addPoints("Webserver (Apache)", 15)
		case strings.Contains(bLower, "nginx"):
			addPoints("Webserver (Nginx)", 15)
		case strings.Contains(bLower, "iis"):
			addPoints("Webserver (Microsoft IIS)", 15)
		case strings.Contains(bLower, "litespeed"):
			addPoints("Webserver (LiteSpeed)", 15)
		case strings.Contains(bLower, "windows"):
			addPoints("Windows Gerät", 10)
		case strings.Contains(bLower, "debian"), strings.Contains(bLower, "ubuntu"), strings.Contains(bLower, "linux"):
			addPoints("Linux/Unix Gerät", 10)
		case strings.Contains(bLower, "openwrt"):
			addPoints("Router (OpenWrt)", 12)
		case strings.Contains(bLower, "router"):
			addPoints("Router Gerät", 10)
		case strings.Contains(bLower, "printer"), strings.Contains(bLower, "hp"), strings.Contains(bLower, "epson"), strings.Contains(bLower, "brother"):
			addPoints("Drucker Gerät", 10)
		case strings.Contains(bLower, "synology"):
			addPoints("Synology NAS", 14)
		case strings.Contains(bLower, "qnap"):
			addPoints("QNAP NAS", 14)
		case strings.Contains(bLower, "raspbian"):
			addPoints("Raspberry Pi", 14)
		}
	}

	// Hostname (niedrig)
	if regexp.MustCompile(`^desktop-`).MatchString(h) {
		addPoints("Windows PC (Hostname)", 6)
	}
	if regexp.MustCompile(`rasp`).MatchString(h) {
		addPoints("Raspberry Pi (Hostname)", 7)
	}
	if regexp.MustCompile(`android`).MatchString(h) {
		addPoints("Android Gerät", 6)
	}
	if regexp.MustCompile(`printer|hp|epson|brother`).MatchString(h) {
		addPoints("Drucker Gerät (Hostname)", 6)
	}
	if regexp.MustCompile(`nas`).MatchString(h) {
		addPoints("NAS (Hostname)", 6)
	}
	if regexp.MustCompile(`router|gateway|tplink|linksys|asus|netgear`).MatchString(h) {
		addPoints("Router (Hostname)", 6)
	}

	// Ergebnis: Höchste Punktzahl
	maxType := "Unbekannt"
	maxPoints := 0
	for deviceType, pts := range scores {
		if pts > maxPoints {
			maxPoints = pts
			maxType = deviceType
		}
	}

	return maxType
}
