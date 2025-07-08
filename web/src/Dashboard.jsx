import { Box, Paper, Typography } from "@mui/material";
import { useEffect, useRef } from "react";
import { Chart, registerables } from "chart.js";
Chart.register(...registerables);

export default function Dashboard({ devices }) {
  const portChartRef = useRef(null);
  const typeChartRef = useRef(null);

  useEffect(() => {
    // Ports zählen
    const portCounts = {};
    devices.forEach((d) =>
      (d.Ports || []).forEach((p) => (portCounts[p] = (portCounts[p] || 0) + 1))
    );
    const portLabels = Object.keys(portCounts);
    const portData = Object.values(portCounts);

    // Gerätetypen zählen
    const typeCounts = {};
    devices.forEach((d) => {
      const t = d.DeviceType || "Unbekannt";
      typeCounts[t] = (typeCounts[t] || 0) + 1;
    });
    const typeLabels = Object.keys(typeCounts);
    const typeData = Object.values(typeCounts);

    // Ports-Chart
    if (portChartRef.current) {
      if (portChartRef.current.chartInstance) {
        portChartRef.current.chartInstance.destroy();
      }
      if (portData.length > 0) {
        portChartRef.current.chartInstance = new Chart(portChartRef.current, {
          type: "bar",
          data: {
            labels: portLabels,
            datasets: [
              {
                label: "Offene Ports",
                data: portData,
                backgroundColor: "rgba(44, 123, 255, 0.85)",
                borderRadius: 6,
              },
            ],
          },
          options: {
            plugins: { legend: { display: false } },
            scales: {
              x: { ticks: { color: "#eee" }, grid: { color: "#333" } },
              y: { ticks: { color: "#eee" }, grid: { color: "#333" } },
            },
            responsive: true,
            maintainAspectRatio: false,
          },
        });
      }
    }

    // Gerätetypen-Chart
    if (typeChartRef.current) {
      if (typeChartRef.current.chartInstance) {
        typeChartRef.current.chartInstance.destroy();
      }
      if (typeData.length > 0) {
        typeChartRef.current.chartInstance = new Chart(typeChartRef.current, {
          type: "pie",
          data: {
            labels: typeLabels,
            datasets: [
              {
                label: "Gerätetypen",
                data: typeData,
                backgroundColor: [
                  "rgba(44,123,255,0.8)",
                  "rgba(0,200,100,0.8)",
                  "rgba(200,0,100,0.8)",
                  "rgba(255,210,40,0.8)",
                ],
              },
            ],
          },
          options: {
            plugins: {
              legend: {
                labels: { color: "#eee", font: { size: 14 } },
              },
            },
            responsive: true,
            maintainAspectRatio: false,
          },
        });
      }
    }

    // CleanUp
    return () => {
      if (portChartRef.current && portChartRef.current.chartInstance) {
        portChartRef.current.chartInstance.destroy();
      }
      if (typeChartRef.current && typeChartRef.current.chartInstance) {
        typeChartRef.current.chartInstance.destroy();
      }
    };
  }, [devices]);

  return (
    <Box
      sx={{
        display: "flex",
        gap: 8,
        flexWrap: "wrap",
        justifyContent: "center",
        width: "100%",
      }}
    >
      <Paper
        elevation={4}
        sx={{
          p: 4,
          flex: "1 1 350px",
          bgcolor: "#1e2127",
          borderRadius: 4,
          maxWidth: 420,
          minWidth: 340,
          height: 320,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography
          variant="h6"
          align="center"
          sx={{ mb: 2, color: "#7ebfff" }}
        >
          Offene Ports
        </Typography>
        <Box sx={{ width: "100%", height: "100%", minHeight: 200 }}>
          <canvas ref={portChartRef} style={{ width: "100%", height: 200 }} />
        </Box>
      </Paper>
      <Paper
        elevation={4}
        sx={{
          p: 4,
          flex: "1 1 350px",
          bgcolor: "#1e2127",
          borderRadius: 4,
          maxWidth: 420,
          minWidth: 340,
          height: 320,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography
          variant="h6"
          align="center"
          sx={{ mb: 2, color: "#7ebfff" }}
        >
          Gerätetypen
        </Typography>
        <Box sx={{ width: "100%", height: "100%", minHeight: 200 }}>
          <canvas ref={typeChartRef} style={{ width: "100%", height: 200 }} />
        </Box>
      </Paper>
    </Box>
  );
}
