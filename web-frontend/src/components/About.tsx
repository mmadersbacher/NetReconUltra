import React from 'react';
import { motion } from 'framer-motion';

export const About: React.FC = () => (
  <motion.section
    id="about"
    className="card"
    initial={{ opacity: 0, y: 40 }}
    animate={{ opacity: 1, y: 0 }}
    transition={{ duration: 0.7, type: 'spring', delay: 0.15 }}
    style={{ margin: '2.5rem auto', textAlign: 'center', maxWidth: 720 }}
  >
    <div style={{ height: 3, width: '100%', background: 'linear-gradient(90deg,#10b6ef,#02e3ff 80%)', borderRadius: '14px 14px 0 0', marginBottom: 12 }} />
    <h2>Über NetReconUltra</h2>
    <p style={{ fontSize: 18, color: 'var(--color-muted)', marginTop: 14 }}>
      NetReconUltra wurde entwickelt, um jedem Nutzer – egal ob Einsteiger oder Fortgeschrittener – eine einfache und visuelle Analyse des eigenen Netzwerks zu ermöglichen.
      <br /><br />
      Keine komplizierte Installation, keine Vorkenntnisse nötig: Einfach öffnen, Scan laden und alles verstehen.
      <br /><br />
      Du willst wissen, was in deinem WLAN los ist? Mit NetReconUltra siehst du alles – sicher, übersichtlich, auf den Punkt.
    </p>
  </motion.section>
);
