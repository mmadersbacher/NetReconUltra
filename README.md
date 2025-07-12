# NetRecon Ultra

**NetRecon Ultra** ist ein hochmodernes, von Grund auf selbst entwickeltes Netzwerk-Discovery- und Analyse-Tool. Es richtet sich an Systemadministratoren, Sicherheitsteams und IT-Profis, die kompromisslose Übersicht und Kontrolle über ihre Netzwerke suchen – ohne Cloud-Abhängigkeiten, ohne Telemetrie, ohne überflüssige Features.

Im Gegensatz zu überladenen Enterprise-Lösungen kombiniert NetRecon Ultra Performance (Go, Goroutines) mit maximaler Transparenz und einem vollständig lokal laufenden Web-Dashboard. Alle Scans werden sauber dokumentiert, historisiert und lassen sich jederzeit analysieren

---

## 🚀 Features

* **Paralleler Netzwerkscan mit Go**
  ICMP-Ping, TCP-Portscan, Hostname-Erkennung, Banner-Grabbing, einfache Betriebssystemeinschätzung via TTL

* **Strukturierte JSON-Reports mit Zeitstempel**
  Alle Scans werden versioniert abgelegt (z. B. `scan_2025-07-01_18-30-00.json`), unabhängig vom Frontend nutzbar

* **Modernes Web-Dashboard mit React + D3.js**
  Geräteübersicht, Filter & Suche, visuelle Topologie-Anzeige, importierbare Scan-Daten

* **Modulare Architektur**
  Der Core ist in Go geschrieben, vollständig erweiterbar um OS-Fingerprinting, ARP/MAC-Vendor-Erkennung, Web-API, IPv6-Support usw.

---

## 🧰 Tech-Stack

```text
Backend:    Go 1.22+       (Konzurrentes Scanning mit Goroutines)
Frontend:   React + Vite   (Modulares Dashboard mit D3.js Visualisierung)
Format:     JSON Reports   (Scan-Daten, historisiert und importierbar)
Visuals:    D3.js, vis.js   (Graphische Geräte-Topologie, dynamisch)
```

---

## 📁 Projektstruktur

```text
.
├── cmd/        # CLI-Startpunkt (main.go)
├── core/       # Scan- & Analyse-Logik (modular)
├── data/       # OUI-Datenbank (MAC → Hersteller)
├── logs/       # Scan-Reports (JSON, Zeitstempel)
├── models/     # Datenstrukturen für Geräte & Reports
├── utils/      # Hilfsfunktionen: Netzwerk, Logging, Parsing
├── web/        # React-Frontend: Dashboard, Graphen, UI
├── go.mod, go.sum
├── LICENSE
└── README.md
```

---

## ⚙️ Installation & Betrieb

### Backend (Go)

```bash
cd cmd
go run main.go
```

* Erkennt automatisch das lokale Subnetz (z. B. 192.168.x.0/24)
* Führt Ping-, Port-, Banner- und Device-Scan durch
* Speichert Report als `logs/scan_YYYY-MM-DD_HH-MM-SS.json`
* Voraussetzung: **Go 1.22+**

### Frontend (React)

```bash
cd web
npm install
npm start
```

* Startet das Dashboard unter [http://localhost:3000](http://localhost:3000)
* Voraussetzung: **Node.js 18+**

---

## 🔄 Scan-Ergebnisse ins Dashboard laden

1. **Report auswählen**:

   ```bash
   ls ../logs/scan_*.json
   ```
2. **Report ins Frontend kopieren**:

   ```bash
   cp ../logs/scan_2025-07-01_18-30-00.json web/logs/latest.json
   ```
3. **Frontend starten und analysieren**

   * Das Dashboard liest immer `web/logs/latest.json`

> ⚠️ Ohne diesen Schritt werden keine Daten angezeigt.

---

## 📈 Typischer Workflow

```text
1. Scan starten         → Backend mit `go run main.go`
2. Ergebnis wählen      → Report im logs/-Verzeichnis auswählen
3. Report kopieren      → in `web/logs/latest.json` einfügen
4. Dashboard starten    → `npm start` im web/-Verzeichnis
5. Analyse durchführen  → Geräte, Ports, Graphen auswerten
```

---

## 🧪 Beispiel-Anwendungsfälle

* Schnelle Inventarisierung eines internen Netzwerks (Home-Office, KMU)
* Aufdeckung offener Ports auf IoT-Geräten oder veralteter Dienste
* Vorbereitung für Pentests oder Schwachstellenbewertungen
* Vergleich von Netzwerkzuständen über die Zeit durch Report-Historie

---

## 🔓 Lizenz

Dieses Projekt steht unter der **MIT-Lizenz**.
Siehe [LICENSE](./LICENSE) für Details.

---

## 📌 Hinweise & Ausblick

* Aktuell keine Live-Kopplung zwischen Backend ↔ Frontend (kein Echtzeit-Trigger)
* IPv6, MAC-Vendor-Erkennung und OS-Fingerprinting sind geplant
* Ziel: Vollständig lokale, professionelle Netzwerk-Analyse ohne Drittanbieter-Abhängigkeiten

---

> Fragen, Ideen oder Feedback?
> [Issue eröffnen](https://github.com/DEIN_USERNAME/NetReconUltra/issues) oder forke das Projekt!
