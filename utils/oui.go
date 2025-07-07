package utils

import (
	"bufio"
	"os"
	"strings"
)

var ouiMap = make(map[string]string)

func LoadOUI(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 8 {
			continue
		}
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) == 2 {
			ouiMap[strings.ToUpper(strings.ReplaceAll(parts[0], ":", ""))] = parts[1]
		}
	}
	return scanner.Err()
}

func LookupVendor(mac string) string {
	key := strings.ToUpper(strings.ReplaceAll(mac, ":", ""))
	if len(key) >= 6 {
		key = key[:6]
	}
	if vendor, ok := ouiMap[key]; ok {
		return vendor
	}
	return "Unbekannt"
}
