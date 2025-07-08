# NetRecon Ultra

[![Go](https://img.shields.io/badge/Go-1.22-blue?logo=go)](https://golang.org/)
[![React](https://img.shields.io/badge/React-18.2-blue?logo=react)](https://react.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> **NetRecon Ultra** ist ein modularer Netzwerk-Scanner und Analyzer für kleine und mittlere Netzwerke.  
> Backend in Go, Frontend in React – dynamisch, erweiterbar und mit Fokus auf Praxistauglichkeit, Performance und Klarheit.

---

## 🔎 Funktionen

- **Paralleler Netzwerkscan:** ICMP, TCP-Portscan, Banner-, Hostname- und Device-Detection per Go Goroutines
- **Gerätetyp-Schätzung:** durch Portmuster, Banner und Hostname-Analyse
- **JSON-Export & Scan-History:** Jeder Scan als Report, automatisch mit Zeitstempel archiviert
- **Interaktives Web-Frontend:** Visualisierung als Dashboard, Geräte-Tabelle, Netzwerkgraph (React + D3.js)
- **Erweiterbar:** Architektur ausgelegt für künftige Features wie OS-Fingerprinting, ARP/MAC/Vendor, Web-API und IPv6

---

## 📁 Projektstruktur

```text
.
├── cmd/
│   └── main.go
├── core/
│   ├── bannergrab.go
│   ├── devicetype.go
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
├── go.mod
├── go.sum
├── LICENSE
├── logs/
│   ├── latest.json
│   ├── scan_2025-07-08_00-14-58.json
│   ├── scan_2025-07-08_01-15-09.json
│   ├── scan_2025-07-08_01-18-26.json
│   ├── scan_2025-07-08_01-21-26.json
│   ├── scan_2025-07-08_01-51-07.json
│   ├── scan_2025-07-08_02-09-32.json
│   ├── scan_2025-07-08_02-21-44.json
│   └── scan_2025-07-08_16-00-34.json
├── models/
│   └── types.go
├── README.md
├── utils/
│   ├── log.go
│   ├── network.go
│   └── oui.go
└── web/
    ├── eslint.config.js
    ├── index.html
    ├── node_modules/
    ├── package.json
    ├── package-lock.json
    ├── public/
    ├── README.md
    ├── src/
    └── vite.config.js

🚀 Schnellstart
Backend (Go)

cd cmd
go run main.go

    Scan-Ergebnisse werden als JSON-Logs im Verzeichnis logs/ gespeichert

Frontend (React)

cd web
npm install
npm start

    Web-Frontend unter http://localhost:3000

    Achtung: Damit die Daten im Dashboard sichtbar sind, muss eine aktuelle Scan-Datei als logs/latest.json im Web-Frontend liegen:

cp ../logs/scan_2025-07-08_16-00-34.json web/logs/latest.json

🖥️ Architektur-Highlights

    Backend (Go):

        CLI-Anwendung zur Subnetz-Erkennung, Host-Discovery, Portscan, Banner- und Gerätetyp-Erkennung

        Modulare Core-Komponenten (scanner, report, osdetect, bannergrab, history)

    Frontend (React):

        Geräteübersicht als dynamische Tabelle mit Filter/Suche

        Netzwerkgraph (D3.js) für Übersicht der Topologie

        Übersichtliche Dashboards, moderne UI-Komponenten (Material UI)

    Datenhaltung:

        Alle Scans als strukturierte JSON-Dateien in /logs/ (inkl. History)

        Webfrontend arbeitet unabhängig und kann beliebige Reports anzeigen

⚠️ Hinweise

    Das Frontend ist derzeit nicht direkt mit dem Backend gekoppelt (kein Live-Trigger).

    Scan-Resultate müssen manuell als logs/latest.json ins Web-Frontend kopiert werden.

    Die Architektur ist bewusst modular, um spätere Features (z. B. ARP/MAC/Vendor, API, OS-Fingerprinting) einfach integrieren zu können.

📄 Lizenz

MIT License – Nutzung, Weiterentwicklung und Integration ausdrücklich erwünscht.
