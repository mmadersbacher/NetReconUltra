package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	ouiCache map[string]string
	once     sync.Once
)

// Lädt die OUI-Tabelle (nur einmal, thread-safe, alles Uppercase)
func loadOUI(ouiFile string) {
	ouiCache = make(map[string]string)
	file, err := os.Open(ouiFile)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 8 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		oui := strings.ToUpper(strings.ReplaceAll(fields[0], ":", ""))
		vendor := strings.Join(fields[1:], " ")
		ouiCache[oui] = vendor
	}
}

// Holt den Vendor zu einer MAC (OUI-File wird beim ersten Call gecached)
func LookupVendor(mac, ouiFile string) string {
	once.Do(func() { loadOUI(ouiFile) })
	if ouiCache == nil {
		fmt.Printf("[DEBUG] OUI-Cache nicht geladen!\n")
		return "Unbekannt"
	}
	key := strings.ToUpper(strings.ReplaceAll(mac, ":", ""))
	if len(key) < 6 {
		fmt.Printf("[DEBUG] MAC zu kurz: %s (Key: %s)\n", mac, key)
		return "Unbekannt"
	}
	oui := key[:6]
	vendor, ok := ouiCache[oui]
	if !ok {
		fmt.Printf("[DEBUG] MAC: %s → OUI: %s → Vendor: Unbekannt\n", mac, oui)
		return "Unbekannt"
	}
	fmt.Printf("[DEBUG] MAC: %s → OUI: %s → Vendor: %s\n", mac, oui, vendor)
	return vendor
}
