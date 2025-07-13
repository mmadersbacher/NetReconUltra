import React, { useEffect, useRef } from 'react';
import { Network } from 'vis-network/standalone';

interface Device {
  id: string;
  label: string;
  ip: string;
}

interface Connection {
  from: string;
  to: string;
}

interface Props {
  devices: Device[];
  connections: Connection[];
}

export const NetworkGraph: React.FC<Props> = ({ devices, connections }) => {
  const containerRef = useRef<HTMLDivElement>(null);
  const networkRef = useRef<Network | null>(null);

  useEffect(() => {
    if (!containerRef.current) return;

    const nodes = devices.map(d => ({ id: d.id, label: `${d.label}\n${d.ip}` }));
    const edges = connections.map(c => ({ from: c.from, to: c.to }));

    if (networkRef.current) {
      networkRef.current.setData({ nodes, edges });
    } else {
      networkRef.current = new Network(containerRef.current, { nodes, edges }, {
        nodes: {
          shape: 'box',
          margin: {
            top: 10,
            right: 10,
            bottom: 10,
            left: 10,
          },
          font: { multi: true, color: '#ffffff' },
          color: {
            background: '#2c3e50',
            border: '#2980b9',
            highlight: { background: '#3498db', border: '#2980b9' }
          }
        },
        edges: {
          arrows: 'to',
          color: '#3498db',
          smooth: true,
        },
        physics: {
          stabilization: false,
          barnesHut: { gravitationalConstant: -2000 },
        },
        interaction: {
          hover: true,
          zoomView: true,
          dragView: true,
        },
        layout: {
          improvedLayout: true,
        },
      });
    }
  }, [devices, connections]);

  return <div ref={containerRef} style={{ height: '400px', border: '1px solid #444', borderRadius: '6px', backgroundColor: '#1e1e2f' }} />;
};
