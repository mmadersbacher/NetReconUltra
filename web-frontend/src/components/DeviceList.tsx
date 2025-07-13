import React from 'react';
import { motion } from 'framer-motion';
import { Monitor, Server, Printer, Globe } from 'lucide-react';

interface Port { Port: number; Open: boolean; Service: string; }
interface Device {
  IP: string;
  Hostname?: string;
  Vendor?: string;
  MAC?: string;
  DeviceType?: string;
  Ports?: Port[] | null;
}

function getIcon(type?: string) {
  if (!type) return <Globe size={18} color="#10b6ef" style={{ marginRight: 6 }} />;
  if (type.toLowerCase().includes('drucker')) return <Printer size={18} color="#10b6ef" style={{ marginRight: 6 }} />;
  if (type.toLowerCase().includes('web')) return <Server size={18} color="#10b6ef" style={{ marginRight: 6 }} />;
  if (type.toLowerCase().includes('windows')) return <Monitor size={18} color="#10b6ef" style={{ marginRight: 6 }} />;
  return <Globe size={18} color="#10b6ef" style={{ marginRight: 6 }} />;
}

export const DeviceList: React.FC<{ devices: Device[] }> = ({ devices }) => (
  <motion.section
    id="devices"
    className="card"
    initial={{ opacity: 0, y: 40 }}
    animate={{ opacity: 1, y: 0 }}
    transition={{ duration: 0.7, type: 'spring', delay: 0.1 }}
  >
    <div style={{ height: 3, width: '100%', background: 'linear-gradient(90deg,#10b6ef,#02e3ff 80%)', borderRadius: '14px 14px 0 0', marginBottom: 12 }} />
    <h2 style={{ marginBottom: 24 }}>Geräteübersicht</h2>
    <div style={{ overflowX: 'auto' }}>
      <table style={{ width: '100%', borderSpacing: 0, fontSize: 15 }}>
        <thead>
          <tr style={{ color: 'var(--color-accent)', background: 'var(--color-panel)' }}>
            <th style={{ padding: 8 }}>IP</th>
            <th style={{ padding: 8 }}>Hostname</th>
            <th style={{ padding: 8 }}>Typ</th>
            <th style={{ padding: 8 }}>Ports (offen/gesamt)</th>
            <th style={{ padding: 8 }}>MAC</th>
            <th style={{ padding: 8 }}>Vendor</th>
          </tr>
        </thead>
        <tbody>
          {devices.map((d, i) => (
            <tr key={i} style={{ borderBottom: '1px solid #223', transition: 'background 0.1s' }}>
              <td style={{ padding: 8 }}>{d.IP}</td>
              <td style={{ padding: 8 }}>{d.Hostname || <span style={{ color: '#555' }}>–</span>}</td>
              <td style={{ padding: 8 }}>{getIcon(d.DeviceType)}{d.DeviceType || <span style={{ color: '#555' }}>–</span>}</td>
              <td style={{ padding: 8 }}>
                {d.Ports ? (
                  <span>
                    <span style={{ color: '#10b6ef', fontWeight: 700 }}>{d.Ports.filter(p => p.Open).length}</span>
                    <span style={{ color: '#9aa4c1' }}>/</span>
                    <span style={{ color: '#fff' }}>{d.Ports.length}</span>
                  </span>
                ) : '–'}
              </td>
              <td style={{ padding: 8 }}>{d.MAC || '–'}</td>
              <td style={{ padding: 8 }}>{d.Vendor || '–'}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  </motion.section>
);
