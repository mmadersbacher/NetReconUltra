package models

import "time"

type PortResult struct {
	Port     int
	Open     bool
	Service  string
	Duration time.Duration
	Banner   string
}

type Device struct {
	IP         string
	MAC        string
	Vendor     string
	Hostname   string
	Ports      []PortResult // Nur models.PortResult!
	Banners    map[string]string
	FoundBy    []string
	DeviceType string
}
