# NetRecon Ultra

NetRecon Ultra ist ein eigenständig entwickeltes Netzwerk-Discovery- und Analyse-Tool für kleine bis mittlere Netzwerke.  
Das Projekt kombiniert ein leistungsstarkes Go-Backend für parallele Netzwerkerkennung mit einem modernen React-Frontend für die Auswertung und Visualisierung der Scanergebnisse.

---

## Inhalt

1. Zielsetzung
2. Features
3. Projektstruktur (Tree)
4. Installation und Betrieb
   - Backend
   - Frontend
   - Kopplung Scan-Log mit Web-UI
5. Nutzung & Workflow
6. Hinweise und Erweiterungen
7. Lizenz & Kontakt

---

## 1. Zielsetzung

NetRecon Ultra wurde entwickelt, um Netzwerke automatisiert und effizient zu scannen, Geräte und Dienste zu erkennen sowie die Ergebnisse nachvollziehbar und professionell aufzubereiten.  
Das Tool richtet sich an Administratoren, IT-Security-Teams und technisch versierte Anwender, die Wert auf Übersicht, Nachvollziehbarkeit und Erweiterbarkeit legen.

---

## 2. Features

- Paralleler Netzwerkscan: ICMP, TCP-Portscan, Banner-Grabbing, Hostname- und Device-Type-Erkennung (Go, Goroutines)
- Export als strukturierte JSON-Reports mit Zeitstempel (Scan-History)
- Modulare Codebasis für Erweiterungen (OS-Fingerprinting, MAC/Vendor, IPv6, Web-API etc.)
- React-Frontend mit Dashboard, Geräte-Tabelle, Filter, Netzwerkgraph (D3.js)
- Reports werden unabhängig vom Frontend gespeichert und können flexibel ausgewertet werden

---

## 3. Projektstruktur

.
├── cmd/ # CLI-Einstiegspunkt (main.go)
├── core/ # Scan- und Analyselogik (Go, modular)
├── data/ # OUI-Datenbank für MAC/Vendor-Erkennung
├── logs/ # Alle Scan-Reports im JSON-Format (inkl. latest.json)
├── models/ # Datentypen für Devices und Reports
├── utils/ # Hilfsfunktionen (Logging, Netzwerktools, OUI-Parsing)
├── web/ # React-Frontend: Dashboard, Visualisierung, Geräteansicht
├── go.mod, go.sum # Go-Abhängigkeiten
├── LICENSE
└── README.md


---

## 4. Installation und Betrieb

### Backend (Go)

```bash
cd cmd
go run main.go

    Erkennt automatisch das lokale Subnetz (manuelle Anpassung möglich)

    Führt Host-Discovery, Portscan, Banner- und Geräteerkennung durch

    Speichert Reports als logs/scan_YYYY-MM-DD_HH-MM-SS.json

    Setzt Go 1.22 oder neuer voraus

Frontend (React)

cd web
npm install
npm start

    Startet ein modernes Dashboard auf http://localhost:3000

    Setzt Node.js 18+ voraus

Kopplung Scan-Log mit Web-UI

Damit das Dashboard Scan-Daten anzeigen kann,
muss eine aktuelle Report-Datei als logs/latest.json ins Web-Verzeichnis kopiert werden.

cp ../logs/scan_YYYY-MM-DD_HH-MM-SS.json logs/latest.json

5. Nutzung & Workflow

    Netzwerk-Scan mit Go-Backend starten (go run main.go im cmd/-Verzeichnis)

    Nach Abschluss gewünschten Report aus /logs/ auswählen

    Report als latest.json ins Web-Frontend kopieren (web/logs/latest.json)

    React-Frontend starten (npm start im web/-Verzeichnis)

    Scanergebnisse im Browser auswerten, filtern und visualisieren

6. Hinweise und Erweiterungen

    Die Kopplung von Backend und Frontend erfolgt derzeit nicht automatisiert (kein Live-Scan-Trigger über das Web-UI).

    Das Projekt ist modular angelegt und kann um zusätzliche Features (z. B. OS-Fingerprinting, ARP/MAC, Web-API, Echtzeit-Visualisierung) erweitert werden.

    Alle Scan-Reports bleiben als History erhalten, was langfristige Netzwerk-Analysen ermöglicht.
