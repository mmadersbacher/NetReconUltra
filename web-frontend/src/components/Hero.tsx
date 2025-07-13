import React from 'react';
import { TechNetBackground } from './TechNetBackground';
import { motion } from "framer-motion";

export const Hero: React.FC = () => (
  <section style={{
    position: 'relative', textAlign: 'center',
    margin: '3.6rem 0 2.2rem 0', padding: '3.5rem 0 2.7rem 0',
    borderRadius: 26, overflow: 'hidden', boxShadow: '0 8px 44px #10b6ef22',
    background: 'linear-gradient(120deg,#151c2fbb,#1b243aee 85%)'
  }}>
    <TechNetBackground />
    <motion.h2
      initial={{ opacity: 0, y: 32 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.7, type: 'spring' }}
      style={{ fontSize: 36, marginBottom: 14, position: 'relative', zIndex: 2 }}
    >
      Volle Übersicht über dein Netzwerk – einfach, schnell, visuell.
    </motion.h2>
    <motion.p
      initial={{ opacity: 0, y: 28 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ delay: .13, duration: .62 }}
      style={{
        fontSize: 21, color: 'var(--color-muted)', maxWidth: 700,
        margin: '0 auto 16px auto', position: 'relative', zIndex: 2
      }}
    >
      NetReconUltra analysiert dein Heimnetzwerk und zeigt dir in Sekunden, welche Geräte verbunden sind, welche Dienste aktiv sind und wie sicher dein Netzwerk ist.<br />
      Egal ob Einsteiger oder Technik-Fan – du bekommst die wichtigsten Infos auf einen Blick, anschaulich und verständlich.
    </motion.p>
    <motion.button
      className="button"
      onClick={() => window.scrollTo({ top: 640, behavior: 'smooth' })}
      initial={{ opacity: 0, scale: 0.98 }}
      animate={{ opacity: 1, scale: 1 }}
      transition={{ delay: .3, duration: .56 }}
      style={{ zIndex: 2, position: 'relative' }}
    >
      Jetzt Netzwerk anzeigen
    </motion.button>
  </section>
);
