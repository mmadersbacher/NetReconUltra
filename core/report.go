package core

import (
	"encoding/json"
	"netreconultra/models"
	"os"
)

func ExportToJSON(devices []models.Device, filename string) error {
	data, err := json.MarshalIndent(devices, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
