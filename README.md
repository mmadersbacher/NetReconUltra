# NetRecon Ultra

NetRecon Ultra ist ein eigenständiges, professionelles Netzwerk-Discovery- und Analyse-Tool für kleine bis mittlere Netzwerke.  
**Backend:** Go (parallele Scans, saubere Reports)  
**Frontend:** React.js (Dashboard, Visualisierung, Filter, Netzwerkgraph)

---

## Inhalt

- [Zielsetzung](#zielsetzung)
- [Features](#features)
- [Projektstruktur](#projektstruktur)
- [Installation & Betrieb](#installation--betrieb)
- [Workflow](#workflow)
- [Hinweise & Erweiterungen](#hinweise--erweiterungen)
- [Lizenz & Kontakt](#lizenz--kontakt)

---

## Zielsetzung

NetRecon Ultra ist gebaut für Admins, Security-Teams und technische Anwender, die  
ein Tool wollen, das *ohne Ballast* Netzwerke scannt, Geräte & Dienste erkennt und die Ergebnisse professionell aufbereitet.  
Keine halbfertigen Scripts, sondern ein solides Werkzeug – modular, nachvollziehbar, erweiterbar.

---

## Features

- **Paralleler Scan:** ICMP (Ping), TCP-Portscan, Banner-Grabbing, Hostname & Device-Type-Erkennung (alles asynchron in Go via Goroutines)
- **JSON-Reports:** Klare, strukturierte Reports mit Zeitstempel, werden automatisch historisiert
- **Modulare Architektur:** Erweiterbar um OS-Fingerprinting, MAC/Vendor, IPv6, Web-API etc.
- **React-Frontend:** Dashboard, Geräte-Tabelle, Filter & Suche, Netzwerkgraph (D3.js), saubere Visualisierung
- **Unabhängige Reports:** Speicherung getrennt vom Frontend, volle Flexibilität zur Auswertung

---

## Projektstruktur

```text
.
├── cmd/        # CLI-Startpunkt (main.go)
├── core/       # Scan-/Analyse-Logik (Go, modular)
├── data/       # OUI-Datenbank für MAC/Vendor-Erkennung
├── logs/       # Alle Scan-Reports (JSON, History)
├── models/     # Datenstrukturen für Devices, Reports
├── utils/      # Hilfsfunktionen (Logging, Netzwerk, OUI-Parsing)
├── web/        # React-Frontend: Dashboard, Visualisierung, Geräteansicht
├── go.mod, go.sum # Go-Abhängigkeiten
├── LICENSE
└── README.md

Installation & Betrieb
Backend (Go)

cd cmd
go run main.go

    Erkennt automatisch das lokale Subnetz (optional konfigurierbar)

    Führt Host-Discovery, Portscan, Banner- und Geräteerkennung durch

    Speichert Reports als logs/scan_YYYY-MM-DD_HH-MM-SS.json

    Voraussetzung: Go 1.22 oder neuer

Frontend (React)

cd web
npm install
npm start

    Startet das Dashboard auf http://localhost:3000

    Voraussetzung: Node.js 18+

Kopplung Scan-Log & Web-UI

Damit das Frontend die Scan-Daten anzeigen kann, muss der Report als
logs/latest.json im Web-Verzeichnis liegen.

cp ../logs/scan_YYYY-MM-DD_HH-MM-SS.json web/logs/latest.json

Workflow

    Netzwerkscan: Backend starten (go run main.go im cmd/-Verzeichnis)

    Report auswählen: Gewünschten Report aus /logs/ kopieren

    Report verlinken:
    cp logs/scan_TIMESTAMP.json web/logs/latest.json

    Frontend starten:
    npm start im web/-Verzeichnis

    Auswertung: Scan im Browser analysieren, filtern, visualisieren

Hinweise & Erweiterungen

    Die Kopplung Backend ↔ Frontend ist aktuell nicht automatisiert (kein Live-Scan/Trigger im Web-UI)

    Das Projekt ist modular: leicht erweiterbar um Features wie OS-Fingerprinting, ARP/MAC, Web-API, Echtzeit-Visualisierung etc.

    Reports werden als History aufbewahrt, für langfristige Netzwerk-Analysen
