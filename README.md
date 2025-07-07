# NetReconUltra

**Modularer, ultraschneller Netzwerk-Scanner in Go.**  
Findet Hosts, Ports, Banner, Hostnamen, Gerätetypen – mit strukturiertem JSON-Export und History/Logs für echte Netzwerkanalyse.

---

## Features

- **Extrem schneller Netzwerkscan** (Parallel, Multithreaded)
- **Host Discovery**: Ping Sweep, Port-Scan
- **Portscan**: vordefinierte und konfigurierbare Ports (TCP)
- **Banner Grabbing**: Liest Service-Banner für HTTP, FTP, Drucker, uvm.
- **Hostname Lookup**: Reverse DNS
- **Gerätetyp- und OS-Erkennung** (per Banner, Ports, Hostname)
- **Strukturierter JSON-Report** (alle Funde, sauber exportiert)
- **Automatisches Log-Verzeichnis**: Reports mit Zeitstempel, `logs/`
- **Delta-/History-Support**: Vergleicht Netzwerkscans im Zeitverlauf (optional)
- **Kompromisslos modularer Code**: Kein File >200 Zeilen, jede Funktion in eigenem Modul

---

## Wie funktioniert's?

1. **Bauen (optional):**
    ```sh
    go build -o netreconultra ./cmd
    ```

2. **Starten (Standard):**
    ```sh
    sudo go run ./cmd scan
    ```
    - Erkennt Subnetz und Interface automatisch

3. **Manuell Subnetz/Interface wählen:**
    ```sh
    sudo go run ./cmd scan <interface> <netzadresse>
    ```
    z.B.  
    ```sh
    sudo go run ./cmd scan wlan0 192.168.8.0
    ```

4. **Reports:**
    - Automatisch nach jedem Scan in `logs/scan_YYYY-MM-DD_HH-MM-SS.json`
    - Immer aktuellster Report als `logs/latest.json`
    - Reports sind **NICHT** im GitHub-Repo!

---

## Beispiel-Output

Starte Netzwerk-Discovery...
Verwende Subnetz: 192.168.8.0/24
Host online (Ping): 192.168.8.133
...
Scan-Ergebnis:
IP Hostname DeviceType Ports Banners FoundBy
192.168.8.133 HPC38E23 Drucker [80 443 8080] {80: HTTP/1.1 505...} [ping portscan]
192.168.8.123 DESKTOP-G2NCQDT Windows PC [139 445] {} [ping portscan]
...
Report gespeichert: logs/scan_2025-07-07_23-31-47.json


---

## Ordnerstruktur

.
├── cmd/ # Entry-Point (main.go)
├── core/ # Hauptmodule (Scanner, Portscan, Ping, Banner, etc)
├── models/ # Datenstrukturen (z.B. Device)
├── utils/ # Helfer (z.B. OUI-Vendor)
├── logs/ # Reports/Exports (nicht im Repo, siehe .gitignore)
├── data/ # Externe Daten (OUI-Liste, Beispiel-Dateien)
├── README.md
├── .gitignore
├── go.mod / go.sum


---

## Warum ist logs/ nicht im Repo?

- Reports sind **lokale Ergebnisse**. Niemand will deine persönlichen Netz-Scans, sondern das Tool.
- `.gitignore` sorgt dafür, dass keine Scans oder sensiblen Daten versehentlich veröffentlicht werden.
- Im Repo ist dafür eine Dummy-Datei `.gitkeep` oder ein `README.md` im Ordner, damit `logs/` immer angelegt ist.

---

## Lizenz

MIT – Siehe LICENSE

---

## Wer sollte das NICHT benutzen?

- Script-Kiddies, die das für illegale Netzwerke nutzen wollen
- Leute, die zu faul sind, README und Doku zu lesen

**Das ist ein Profi-Tool. Wer das nicht versteht, soll Nmap benutzen.**

---

## Autor

Mario Madersbacher  
[mmadersbacher auf GitHub](https://github.com/mmadersbacher)

---
