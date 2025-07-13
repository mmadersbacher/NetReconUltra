import React, { useEffect, useState } from 'react';
import { MainLayout } from './layout/MainLayout';
import { Hero } from './components/Hero';
import { Dashboard } from './components/Dashboard';
import { DeviceList } from './components/DeviceList';
import { About } from './components/About';
// ThemeSwitch ist bereits im Header eingebunden! Keine doppelte Integration nötig.

interface Port {
  Port: number;
  Open: boolean;
  Service: string;
}
interface Device {
  IP: string;
  Hostname?: string;
  Vendor?: string;
  MAC?: string;
  DeviceType?: string;
  Ports?: Port[] | null;
}
interface ScanData { devices: Device[]; }

const calculateStats = (devices: Device[]) => {
  let totalPortsScanned = 0, openPorts = 0;
  devices.forEach(device => {
    if (device.Ports) {
      totalPortsScanned += device.Ports.length;
      openPorts += device.Ports.filter(p => p.Open).length;
    }
  });
  return { activeDevices: devices.length, totalPortsScanned, openPorts };
};

const App: React.FC = () => {
  const [scanData, setScanData] = useState<ScanData | null>(null);
  const [stats, setStats] = useState<{ activeDevices: number; totalPortsScanned: number; openPorts: number } | null>(null);

  // Theme Sync: Hole Theme aus LocalStorage (Initialisierung) – ist aber in ThemeSwitch abgesichert!
  useEffect(() => {
    const theme = (localStorage.getItem('theme') as 'dark' | 'light') || 'dark';
    document.documentElement.setAttribute('data-theme', theme);
  }, []);

  // Laden der Scandaten
  useEffect(() => {
    fetch('/logs/latest.json')
      .then(res => res.json())
      .then(data => {
        setScanData(data);
        setStats(calculateStats(data.devices));
      });
  }, []);

  if (!scanData || !stats) {
    return (
      <div style={{
        textAlign: 'center', marginTop: 80,
        color: 'var(--color-accent)', fontSize: 22, fontFamily: 'Fira Mono, monospace'
      }}>
        <span>Lade Scan-Daten...</span>
      </div>
    );
  }

  return (
    <MainLayout>
      <Hero />
      <Dashboard devices={scanData.devices} stats={stats} />
      <DeviceList devices={scanData.devices} />
      <About />
    </MainLayout>
  );
};

export default App;
