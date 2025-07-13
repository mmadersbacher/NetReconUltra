# <img src="./web-frontend/public/assets/logo.svg" height="36" alt="NetReconUltra Logo" align="left"/> NetReconUltra

[![Go Version](https://img.shields.io/badge/go-%3E=1.22-blue?logo=go)](https://golang.org)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)]()
[![License](https://img.shields.io/github/license/mmadersbacher/NetReconUltra)](LICENSE)
[![Issues](https://img.shields.io/github/issues/mmadersbacher/NetReconUltra?color=blue)](https://github.com/mmadersbacher/NetReconUltra/issues)
[![Frontend Demo](https://img.shields.io/badge/frontend-live-blue?logo=react)](https://mmadersbacher.github.io/NetReconUltra/)

<br/>

**NetReconUltra** ist ein hochperformanter, modularer Netzwerk-Scanner und Visualizer auf ELITE-Niveau.  
Entwickelt für fortgeschrittene Netzwerkerkennung, Geräte- und Dienstanalyse, Fingerprinting und exaktes Reporting –  
visuell, modern, modular.

**Backend:** Go – ultraschneller Scanner, cleane JSON-Reports  
**Frontend:** React + TypeScript – State-of-the-Art UI, animiertes Dashboard

---

## Features

- Automatische Netzwerkerkennung – findet und scannt alle aktiven Subnetze
- Paralleler Ping-/Portscan – blitzschnell, hochskalierbar
- Service-Banner-Grabbing – erkennt laufende Dienste (HTTP/S, FTP, Drucker, SMB, etc.)
- Geräte- und OS-Erkennung – per TTL, Banner, Hostname, Port-Profiling
- Topologische Visualisierung – Übersicht über alle Geräte & Verbindungen (Frontend)
- Live-Statistiken & Dashboard – Offene Ports, Gerätetypen, Services
- History- und Delta-Reports – Vergleich von Scans über Zeit (in Planung)
- Keine sensiblen Daten im Repo: Logs/Reports bleiben ausschließlich lokal
- Voll responsives Elite-UI – ThemeSwitch (Dark/Light), Neon-Canvas, Animationen

---

## Screenshots

| Dashboard (Topologie + Statistiken) | Geräteübersicht (Tabelle) |
|-------------------------------------|--------------------------|
| ![Dashboard Screenshot](./web-frontend/public/assets/screenshot-dashboard.png)<br><sub>Zeigt Topologie-Graph, animierten Hintergrund, StatsChart, ThemeSwitch oben rechts</sub> | ![Devices Screenshot](./web-frontend/public/assets/screenshot-devices.png)<br><sub>Geräte-Tabelle, Icons, offene Ports farbig hervorgehoben</sub> |

*Screenshots sind Platzhalter und können durch eigene Bilder ersetzt werden.*

---

## Installation (Backend & Frontend)

### Backend (Go-Scanner)

```bash
git clone https://github.com/mmadersbacher/NetReconUltra.git
cd NetReconUltra
go mod tidy
go build -o netreconultra ./cmd

Frontend (Visualizer / Dashboard)

cd web-frontend
npm install
npm run dev

    Die Web-Oberfläche läuft dann auf http://localhost:5173

    Die Scan-Reports müssen als /public/logs/latest.json vorliegen (siehe Abschnitt "Reporting").

Benutzung / Quickstart
Netzwerkscan starten

# Automatische Interface-/Subnetz-Erkennung:
sudo ./netreconultra scan

# Oder manuell, z.B.:
sudo ./netreconultra scan wlan0 192.168.8.0

    Der aktuelle Report liegt immer als logs/latest.json

    Alte Reports werden mit Zeitstempel archiviert: logs/scan_YYYY-MM-DD_HH-MM-SS.json

Frontend starten & Scan visualisieren

    Frontend starten (npm run dev)

    Gewünschten Scan-Report als web-frontend/public/logs/latest.json speichern

    Website neu laden – du siehst die komplette Netzwerk-Visualisierung

Projektstruktur (Kurzüberblick)

NetReconUltra/
├── cmd/        # Go-CLI-Entrypoint
├── core/       # Scan-/Analyse-Module (Go)
├── models/     # Typen, Datenstrukturen
├── data/       # MAC/OUI Datenbank
├── logs/       # Alle Reports (JSON, .gitignored)
├── web-frontend/
│   ├── public/
│   │   ├── assets/   # Logos, Screenshots
│   │   └── logs/
│   ├── src/
│   │   ├── components/
│   │   ├── layout/
│   │   └── App.tsx etc.
│   ├── index.css
│   ├── package.json
│   └── ...
├── go.mod / go.sum
├── README.md
└── .gitignore

Reporting, Datenschutz & Logs

    Alle Reports werden ausschließlich lokal gespeichert (logs/)

    Keine gescannten Daten oder privaten Infos verlassen den Rechner

    Der Ordner logs/ ist im Git ignoriert

Tech Stack

Backend:
Go (>=1.22), Standard Library

Frontend:
React 18, TypeScript, Vite
Framer Motion (Animationen)
Lucide Icons
Chart.js (Statistiken)
Custom Canvas (Hero-Background)
Modernes CSS, Neon-Theme, ThemeSwitch
Demo & Live-Preview

Live-Frontend (Demo auf GitHub Pages):
https://mmadersbacher.github.io/NetReconUltra/
(Demo zeigt Beispieldaten – keine Live-Scans im Browser!)
Lizenz & Credits

MIT License – siehe LICENSE
Created by Mario Madersbacher, 2025
Kontakt & Feedback

Fragen, Bugreports, Featurewünsche: GitHub Issues
Portfolio & Kontakt: github.com/mmadersbacher
