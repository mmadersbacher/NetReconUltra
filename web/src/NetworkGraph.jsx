import React, { useEffect, useRef } from "react";
import { Paper, Typography } from "@mui/material";
import { Network } from "vis-network/peer";
import "vis-network/styles/vis-network.css";

function buildGraph(devices) {
  // Nodes: alle Geräte
  const nodes = devices.map((dev, idx) => ({
    id: idx + 1,
    label: `${dev.Hostname || dev.IP}`,
    title: `${dev.IP}\n${dev.DeviceType}`,
    shape: dev.DeviceType && dev.DeviceType.toLowerCase().includes("router")
      ? "box"
      : dev.DeviceType && dev.DeviceType.toLowerCase().includes("drucker")
      ? "triangle"
      : "ellipse",
    color: dev.DeviceType && dev.DeviceType.toLowerCase().includes("router")
      ? "#ffdb3a"
      : dev.DeviceType && dev.DeviceType.toLowerCase().includes("drucker")
      ? "#2affd5"
      : "#7ebfff",
    font: { color: "#222", face: "monospace", size: 17 },
    borderWidth: dev.DeviceType && dev.DeviceType.toLowerCase().includes("router") ? 5 : 2,
  }));

  // Edges: alles was mit Gateway verbunden ist = direkte Verbindung
  const gatewayIndex = devices.findIndex(
    (d) => d.Hostname && d.Hostname.toLowerCase().includes("gateway")
  );
  const edges = [];
  if (gatewayIndex !== -1) {
    devices.forEach((_, idx) => {
      if (idx !== gatewayIndex) {
        edges.push({ from: gatewayIndex + 1, to: idx + 1, color: "#888" });
      }
    });
  }
  return { nodes, edges };
}

export default function NetworkGraph({ devices }) {
  const containerRef = useRef();

  useEffect(() => {
    if (!devices.length) return;

    const { nodes, edges } = buildGraph(devices);

    const net = new Network(containerRef.current, { nodes, edges }, {
      physics: { stabilization: true, barnesHut: { gravitationalConstant: -2500 } },
      layout: { improvedLayout: true },
      nodes: {
        borderWidth: 2,
        color: {
          background: "#23272f",
          border: "#7ebfff",
          highlight: "#459aff",
          hover: "#a3dbff"
        }
      },
      edges: {
        width: 2,
        color: "#88a",
        smooth: { enabled: true, type: "continuous" }
      },
      interaction: {
        hover: true,
        tooltipDelay: 100,
        navigationButtons: true,
        keyboard: true
      }
    });

    // Cleanup
    return () => net && net.destroy();
  }, [devices]);

  return (
    <Paper
      elevation={4}
      sx={{
        bgcolor: "#23272f",
        borderRadius: 4,
        mt: 3,
        p: 2,
        width: "100%",
        height: 420,
        minHeight: 200,
        display: "flex",
        flexDirection: "column"
      }}
    >
      <Typography
        variant="h6"
        align="center"
        sx={{ color: "#7ebfff", mb: 2 }}
      >
        Netzwerk-Topologie
      </Typography>
      <div ref={containerRef} style={{ width: "100%", flex: 1, minHeight: 320 }} />
    </Paper>
  );
}
