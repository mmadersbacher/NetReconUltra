package main

import (
	"fmt"
	"netreconultra/core"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "scan":
		core.RunScan()
	case "report":
		// Platzhalter: Implementier das später, oder schreib klar raus:
		fmt.Println("Report-Feature noch nicht implementiert.")
		// core.GenerateReport() // (wenn fertig)
	case "history":
		fmt.Println("History-Feature noch nicht implementiert.")
		// core.ShowHistory() // (wenn fertig)
	case "help":
		printHelp()
	default:
		fmt.Println("Unbekannter Befehl:", command)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("NetRecon Ultra – Next Level Network Scanner")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  netreconultra scan      # Führt einen Netzwerkscan durch")
	fmt.Println("  netreconultra report    # Generiert einen Report aus den Daten (bald)")
	fmt.Println("  netreconultra history   # Zeigt Scan-History und Vergleiche (bald)")
	fmt.Println("  netreconultra help      # Zeigt diese Hilfe an")
}
