# NetReconUltra

[![Go Version](https://img.shields.io/badge/go-%3E=1.22-blue?logo=go)](https://golang.org)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)]()
[![License](https://img.shields.io/github/license/mmadersbacher/NetReconUltra)](LICENSE)
[![Issues](https://img.shields.io/github/issues/mmadersbacher/NetReconUltra?color=blue)](https://github.com/mmadersbacher/NetReconUltra/issues)

**NetReconUltra** ist ein hochperformanter, modularer Netzwerk-Scanner in Go.  
Entwickelt für fortgeschrittene Netzwerkerkennung, Geräteidentifikation, Service-Fingerprinting und exaktes Reporting im Business- und Security-Kontext.

---

## Features

- Automatische Erkennung von aktivem Interface und Subnetz
- Paralleler Ping- und Portscan mit minimaler Latenz
- Service-Banner-Grabbing (HTTP, FTP, Drucker, uvm.)
- Gerätetyp- und OS-Identifikation (Banner, Ports, Hostname, Fingerprinting)
- Strukturierte JSON-Reports mit Zeitstempel und History/Delta-Support
- Strikte Modularisierung: kein File >200 Zeilen, klar getrennte Komponenten
- Keine sensiblen Daten im Repo: Logs und Reports bleiben ausschließlich lokal

---

## Quickstart

```sh
git clone https://github.com/mmadersbacher/NetReconUltra.git
cd NetReconUltra
go mod tidy

# Build (optional)
go build -o netreconultra ./cmd

# Direkt ausführen (automatische Subnetzerkennung)
sudo go run ./cmd scan

# Manuelles Interface/Subnetz
sudo go run ./cmd scan wlan0 192.168.8.0

Reports werden automatisch in logs/ mit Zeitstempel exportiert.
Der aktuellste Report ist immer unter logs/latest.json verfügbar.
Beispielausgabe

Starte Netzwerk-Discovery...
Verwende Subnetz: 192.168.8.0/24
Host online (Ping): 192.168.8.133
Scan-Ergebnis:
IP              Hostname        DeviceType     Ports          Banners                   FoundBy
192.168.8.133   HPC38E23        Drucker        [80 443 8080]  {80:HTTP/1.1 ...}         [ping portscan]
192.168.8.123   DESKTOP-G2NCQDT Windows PC     [139 445]      {}                        [ping portscan]
...
Report gespeichert: logs/scan_2025-07-07_23-31-47.json

Ordnerstruktur

NetReconUltra/
├── cmd/
│   └── main.go
├── core/
│   ├── bannergrab.go
│   ├── history.go
│   ├── hostdiscovery.go
│   ├── osdetect.go
│   ├── pingsweep.go
│   ├── portscan.go
│   ├── ports.go
│   ├── report.go
│   └── scanner.go
├── data/
│   └── oui.txt
├── logs/
│   └── .gitkeep
├── models/
│   └── types.go
├── utils/
│   ├── export.go
│   ├── log.go
│   └── network.go
├── README.md
├── .gitignore
├── go.mod
├── go.sum

Reporting, Logs und Datenschutz

    Logs und Reports werden ausschließlich lokal im Verzeichnis logs/ gespeichert.

    Der Ordner logs/ ist nicht Teil des GitHub-Repos (siehe .gitignore).

    Keine gescannten Daten oder privaten Netzwerkinfos verlassen den Rechner.

Lizenz

MIT – siehe LICENSE.
Kontakt / Autor

Mario Madersbacher
GitHub Profil
