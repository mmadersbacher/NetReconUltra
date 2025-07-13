import React from 'react';
import { Bar } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

interface Stats {
  activeDevices: number;
  totalPortsScanned: number;
  openPorts: number;
}

interface Props {
  stats: Stats;
}

export const StatsChart: React.FC<Props> = ({ stats }) => {
  const data = {
    labels: ['Aktive Geräte', 'Gescannt Ports', 'Offene Ports'],
    datasets: [
      {
        label: 'NetRecon Ultra',
        data: [stats.activeDevices, stats.totalPortsScanned, stats.openPorts],
        backgroundColor: ['#3498db', '#f1c40f', '#2ecc71'],
        borderRadius: 4,
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top' as const,
        labels: { color: '#e0e0e0' },
      },
      title: {
        display: true,
        text: 'Scan Statistiken',
        color: '#e0e0e0',
        font: { size: 18 },
      },
      tooltip: {
        enabled: true,
      },
    },
    scales: {
      x: { ticks: { color: '#e0e0e0' } },
      y: { ticks: { color: '#e0e0e0' }, beginAtZero: true },
    },
  };

  return <Bar data={data} options={options} />;
};
