 NetRecon Ultra

**NetRecon Ultra** ist ein modular aufgebautes Netzwerk-Discovery- und Analyse-Tool für kleine bis mittlere Netzwerke, entwickelt in Go (Backend) und React (Frontend).  
Das Tool wurde für praxisnahe, schnelle Netzwerkerkennung, Host- und Dienstidentifikation sowie professionelle Auswertung und Visualisierung konzipiert.

---

## Funktionen

- Paralleler Netzwerk-Scan (Ping, Portscan, Banner, Hostname, Device-Typ-Erkennung)
- JSON-Report mit Zeitstempel für jeden Scanlauf, sowie fortlaufende History
- Export der Scan-Ergebnisse in strukturierte Logs zur weiteren Auswertung
- React-basiertes Web-Frontend für interaktive Geräteübersicht und Netzwerkgraph
- Erweiterbar um OS-Fingerprinting, MAC/Vendor-Erkennung, Web-API, IPv6-Unterstützung

---

## Verzeichnisstruktur

.
├── cmd/ # Einstiegspunkt für CLI (main.go)
├── core/ # Hauptlogik für Scan, Discovery, Banner, OS-Detection etc.
├── data/ # Zusatzdaten wie OUI-Datenbank für MAC/Vendor
├── logs/ # Alle Scan-Reports im JSON-Format (inkl. latest.json)
├── models/ # Datentypen für Devices und Reports
├── utils/ # Hilfsfunktionen (Logging, Netzwerktools, OUI-Parsing)
├── web/ # React-Frontend: Dashboard, Visualisierung, Geräteansicht
├── go.mod, go.sum # Go-Abhängigkeiten
├── LICENSE
└── README.md


---

## Installation & Nutzung

### 1. Backend (Go)

```bash
cd cmd
go run main.go

    Das Backend erkennt das lokale Subnetz, führt Discovery & Scans durch und speichert Reports in /logs.

2. Web-Frontend (React)

cd web
npm install
npm start

    Das Frontend ist standardmäßig auf http://localhost:3000 erreichbar.

    Hinweis:
    Für die Anzeige von Scan-Ergebnissen muss eine Datei latest.json im Verzeichnis web/logs/ liegen.
    Diese Datei kann aus einem beliebigen Scan aus /logs/ kopiert werden.

cp ../logs/scan_YYYY-MM-DD_HH-MM-SS.json logs/latest.json

Hinweise

    Das Frontend ist derzeit nicht direkt mit dem Backend verbunden (kein Live-Scan-Trigger über das Web-UI).

    Die Codebasis ist modular ausgelegt und ermöglicht die Erweiterung um neue Funktionen (z. B. erweiterte OS-Erkennung, ARP/MAC-Analyse, Web-API).

    Reports werden standardmäßig in /logs/ abgelegt und können für langfristige Auswertungen genutzt werden.

Lizenz

Dieses Projekt steht unter der MIT-Lizenz.
Kontakt

Fragen oder Interesse an einer technischen Zusammenarbeit?
Kontakt: mario.madersbacher.2008@gmail.com
