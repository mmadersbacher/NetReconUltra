# NetRecon Ultra

NetRecon Ultra ist ein modular aufgebautes Netzwerk-Discovery- und Analyse-Tool für kleine bis mittlere Netzwerke.  
Backend in Go, Frontend in React. Ziel ist eine schnelle, nachvollziehbare Netzwerkerkennung und eine strukturierte, interaktive Auswertung über ein Webfrontend.

## Features

- Paralleler Netzwerkscan: ICMP, TCP-Portscan, Banner-, Hostname- und Device-Typ-Erkennung (Go)
- JSON-Export und Scan-History: Jeder Scan als Report, automatisch mit Zeitstempel archiviert
- Web-Frontend: Dashboard, Geräte-Tabelle, Netzwerkgraph (React, D3.js, Material UI)
- Architektur vorbereitet für OS-Fingerprinting, MAC/Vendor, API, IPv6

## Projektstruktur

.
├── cmd/
│ └── main.go
├── core/
│ ├── bannergrab.go
│ ├── devicetype.go
│ ├── history.go
│ ├── hostdiscovery.go
│ ├── osdetect.go
│ ├── pingsweep.go
│ ├── portscan.go
│ ├── ports.go
│ ├── report.go
│ └── scanner.go
├── data/
│ └── oui.txt
├── go.mod
├── go.sum
├── LICENSE
├── logs/
│ ├── latest.json
│ ├── scan_2025-07-08_00-14-58.json
│ ├── scan_2025-07-08_01-15-09.json
│ ├── scan_2025-07-08_01-18-26.json
│ ├── scan_2025-07-08_01-21-26.json
│ ├── scan_2025-07-08_01-51-07.json
│ ├── scan_2025-07-08_02-09-32.json
│ ├── scan_2025-07-08_02-21-44.json
│ └── scan_2025-07-08_16-00-34.json
├── models/
│ └── types.go
├── README.md
├── utils/
│ ├── log.go
│ ├── network.go
│ └── oui.go
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


## Installation und Nutzung

Backend (Go):

```bash
cd cmd
go run main.go

Scan-Ergebnisse werden als JSON-Logs im Verzeichnis logs/ gespeichert.

Frontend (React):

cd web
npm install
npm start

Das Web-Frontend ist standardmäßig erreichbar unter http://localhost:3000.

Damit die Scandaten im Frontend angezeigt werden, muss eine aktuelle Scan-Datei als logs/latest.json im Web-Frontend liegen.
Beispiel:

cp ../logs/scan_2025-07-08_16-00-34.json logs/latest.json

Hinweise

    Das Frontend ist derzeit nicht direkt mit dem Backend gekoppelt (kein Live-Scan-Trigger).

    Scan-Resultate müssen manuell als logs/latest.json ins Web-Frontend kopiert werden.

    Die Architektur ist modular und für zukünftige Erweiterungen vorbereitet.
