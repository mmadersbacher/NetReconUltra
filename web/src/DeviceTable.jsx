import { useState } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Chip,
  TextField,
  Box,
} from "@mui/material";

export default function DeviceTable({ devices }) {
  const [query, setQuery] = useState("");
  const filtered = devices.filter((d) =>
    [d.IP, d.Hostname, d.DeviceType, (d.Ports || []).join(","), (d.FoundBy || []).join(",")]
      .join(" ")
      .toLowerCase()
      .includes(query.toLowerCase())
  );

  return (
    <Paper
      elevation={4}
      sx={{
        p: 3,
        bgcolor: "#1a1c20",
        borderRadius: 4,
        width: "100%",
        mb: 5,
      }}
    >
      <TextField
        label="Suche"
        variant="outlined"
        size="small"
        fullWidth
        sx={{ mb: 3, background: "#23272e", input: { color: "#fff" }, label: { color: "#ccc" } }}
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        InputProps={{ style: { color: "#fff" } }}
      />
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell sx={{ color: "#7ebfff", fontWeight: "bold" }}>IP</TableCell>
              <TableCell sx={{ color: "#7ebfff", fontWeight: "bold" }}>Hostname</TableCell>
              <TableCell sx={{ color: "#7ebfff", fontWeight: "bold" }}>Gerätetyp</TableCell>
              <TableCell sx={{ color: "#7ebfff", fontWeight: "bold" }}>Ports</TableCell>
              <TableCell sx={{ color: "#7ebfff", fontWeight: "bold" }}>Banner</TableCell>
              <TableCell sx={{ color: "#7ebfff", fontWeight: "bold" }}>FoundBy</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {filtered.map((device, idx) => (
              <TableRow
                key={device.IP}
                sx={{
                  backgroundColor: idx % 2 === 0 ? "#23272e" : "#191b1e",
                }}
              >
                <TableCell sx={{ color: "#fff" }}>{device.IP}</TableCell>
                <TableCell sx={{ color: "#fff" }}>{device.Hostname}</TableCell>
                <TableCell sx={{ color: "#fff" }}>
                  {device.DeviceType || "Unbekannt"}
                </TableCell>
                <TableCell>
                  {(device.Ports || []).map((port) => (
                    <Chip
                      key={port}
                      label={port}
                      size="small"
                      color="primary"
                      sx={{ mr: 0.5, bgcolor: "#2d7fff", color: "#fff" }}
                    />
                  ))}
                </TableCell>
                <TableCell sx={{ color: "#c3c6ca" }}>
                  {device.Banners &&
                    Object.entries(device.Banners).map(([port, banner]) => (
                      <span key={port} style={{ display: "block", fontSize: 12 }}>
                        <b>[{port}]</b> {banner}
                      </span>
                    ))}
                </TableCell>
                <TableCell>
                  {(device.FoundBy || []).map((m) => (
                    <Chip
                      key={m}
                      label={m}
                      size="small"
                      color="secondary"
                      sx={{ mr: 0.5, bgcolor: "#8e44ad", color: "#fff" }}
                    />
                  ))}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Paper>
  );
}
