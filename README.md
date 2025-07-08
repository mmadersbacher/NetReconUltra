NetRecon Ultra

NetRecon Ultra ist ein eigenständig entwickeltes, professionelles Netzwerk-Discovery- und Analyse-Tool für kleine bis mittlere Netzwerke.
Ziel ist es, Admins, Security-Teams und technisch versierten Anwendern endlich einen schnellen, modularen und nachvollziehbaren Netzwerküberblick zu geben – ohne Marketing-Blabla oder Feature-Bloat.
Features

    Paralleler Netzwerkscan (Go, Goroutines): ICMP, TCP-Portscan, Banner-Grabbing, Hostname- und Device-Type-Erkennung

    JSON-Reports mit History: Ergebnisse werden sauber mit Zeitstempel abgelegt

    Professionelles React-Frontend: Dashboard, Geräte-Tabelle, Filter, Suche, Netzwerk-Graph (D3.js)

    Unabhängige Reports: Reports sind unabhängig vom Frontend, können jederzeit importiert und ausgewertet werden

    Modular und erweiterbar: Architektur kann um OS-Fingerprinting, MAC/Vendor, IPv6, Web-API, Echtzeit-Visualisierung usw. erweitert werden

Projektstruktur

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

Um Scan-Ergebnisse im Dashboard anzuzeigen, Report ins Frontend kopieren:

cp ../logs/scan_YYYY-MM-DD_HH-MM-SS.json web/logs/latest.json

    Das Frontend liest Scan-Daten immer aus web/logs/latest.json

    Ohne diesen Schritt werden keine Scan-Daten angezeigt

Workflow

    Netzwerkscan: Backend starten (go run main.go im cmd/-Verzeichnis)

    Report auswählen: Gewünschten Report aus /logs/ wählen

    Report kopieren: Nach web/logs/latest.json verschieben

    Frontend starten: Im web/-Verzeichnis mit npm start das Dashboard öffnen

    Analyse: Scan im Browser auswerten, filtern, visualisieren

Hinweise & Erweiterungen

    Kopplung Backend ↔ Frontend ist aktuell nicht automatisiert (kein Live-Scan/Trigger im Web-UI)

    Das Projekt ist modular und kann einfach um OS-Fingerprinting, ARP/MAC, Web-API, Echtzeit-Visualisierung usw. erweitert werden

    Reports werden historisiert: Jeder Scan bleibt als History erhalten, perfekt für langfristige Analysen
