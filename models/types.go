package models

type Device struct {
	IP       string
	MAC      string
	Vendor   string
	Hostname string
	Ports    []int
	Banners  map[int]string
	FoundBy  []string // z.B. "ping", "portscan", "arp"
}
