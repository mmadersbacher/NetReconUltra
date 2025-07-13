import React from 'react';
import { motion } from 'framer-motion';
import { NetworkGraph } from './NetworkGraph';
import { StatsChart } from './StatsChart';

interface Props {
  devices: any[];
  stats: { activeDevices: number; totalPortsScanned: number; openPorts: number };
}

export const Dashboard: React.FC<Props> = ({ devices, stats }) => {
  const networkDevices = devices.map((d, idx) => ({
    id: String(idx + 1),
    label: d.Hostname || d.DeviceType || d.IP,
    ip: d.IP,
  }));

  const networkConnections: { from: string; to: string }[] = [];

  return (
    <motion.section
      id="dashboard"
      className="card"
      initial={{ opacity: 0, y: 40, scale: 0.98 }}
      animate={{ opacity: 1, y: 0, scale: 1 }}
      transition={{ duration: 0.7, type: 'spring' }}
      style={{ marginTop: 0 }}
    >
      <div style={{ height: 3, width: '100%', background: 'linear-gradient(90deg,#10b6ef,#02e3ff 80%)', borderRadius: '14px 14px 0 0', marginBottom: 12 }} />
      <h2 style={{ textAlign: 'center', marginBottom: 32, fontSize: 24 }}>Netzwerk-Topologie & Statistiken</h2>
      <div className="grid">
        <motion.div initial={{ opacity: 0 }} animate={{ opacity: 1 }} transition={{ delay: 0.2 }}>
          <NetworkGraph devices={networkDevices} connections={networkConnections} />
        </motion.div>
        <motion.div initial={{ opacity: 0 }} animate={{ opacity: 1 }} transition={{ delay: 0.3 }}>
          <StatsChart stats={stats} />
        </motion.div>
      </div>
    </motion.section>
  );
};
