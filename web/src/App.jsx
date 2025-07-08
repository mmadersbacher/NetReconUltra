import * as React from "react";
import { useEffect, useState } from "react";
import { Typography, Box, CssBaseline } from "@mui/material";
import Dashboard from "./Dashboard";
import NetworkGraph from "./NetworkGraph";
import DeviceTable from "./DeviceTable";

export default function App() {
  const [devices, setDevices] = useState([]);

  useEffect(() => {
    fetch("/logs/latest.json")
      .then((r) => r.json())
      .then(setDevices)
      .catch(() => setDevices([]));
  }, []);

  return (
    <Box sx={{ bgcolor: "#15171A", minHeight: "100vh", width: "100vw" }}>
      <CssBaseline />
      <Box
        sx={{
          width: "95vw",
          maxWidth: 1450,
          mx: "auto",
          py: 7,
          display: "flex",
          flexDirection: "column",
          alignItems: "center"
        }}
      >
        <Typography
          variant="h2"
          align="center"
          gutterBottom
          sx={{ mb: 8, color: "#fff", fontWeight: 700 }}
        >
          NetReconUltra – Netzwerk Scan Report
        </Typography>
        <Box
          sx={{
            display: "flex",
            gap: 5,
            flexWrap: "wrap",
            mb: 8,
            justifyContent: "center",
            width: "100%",
          }}
        >
          <Dashboard devices={devices} />
        </Box>
        <Box sx={{ my: 8, width: "100%" }}>
          <NetworkGraph devices={devices} />
        </Box>
        <Box sx={{ mt: 8, width: "100%" }}>
          <Typography
            variant="h5"
            gutterBottom
            align="center"
            sx={{ mb: 4, color: "#c3c6ca" }}
          >
            Geräte im Netzwerk
          </Typography>
          <DeviceTable devices={devices} />
        </Box>
      </Box>
    </Box>
  );
}
