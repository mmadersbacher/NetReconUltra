package models

type Device struct {
	IP         string            `json:"IP"`
	MAC        string            `json:"MAC"`
	Vendor     string            `json:"Vendor"`
	Hostname   string            `json:"Hostname"`
	Ports      []int             `json:"Ports"`
	Banners    map[string]string `json:"Banners"`
	FoundBy    []string          `json:"FoundBy"`
	DeviceType string            `json:"DeviceType"`
}
