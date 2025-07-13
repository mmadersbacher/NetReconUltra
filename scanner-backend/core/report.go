package core

import (
	"encoding/json"
	"netreconultra/models"
	"os"
	"time"
)

// Exportiert ein sauberes JSON-Objekt mit scanDate und devices
func ExportToJSON(devices []models.Device, filename string) error {
	output := struct {
		ScanDate string          `json:"scanDate"`
		Devices  []models.Device `json:"devices"`
	}{
		ScanDate: time.Now().Format("2006-01-02 15:04:05"),
		Devices:  devices,
	}
	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
