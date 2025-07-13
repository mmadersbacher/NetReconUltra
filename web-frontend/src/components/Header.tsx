import React from 'react';
import logo from '../assets/LogoNetRecon_Trans.png';
import { ThemeSwitch } from './ThemeSwitch';

export const Header: React.FC = () => (
  <header style={{
    display: 'flex', alignItems: 'center', gap: 18,
    background: 'var(--color-panel)', padding: '1.15rem 2.2rem',
    borderRadius: '0 0 24px 24px', boxShadow: '0 2px 16px #00e9ff33', marginBottom: 38
  }}>
    <img src={logo} alt="Logo" style={{ height: 46, filter: 'drop-shadow(0 0 5px #10b6ef)' }} />
    <h1 style={{
      fontSize: 30, margin: 0, letterSpacing: 1.2,
      color: 'var(--color-accent2)', textShadow: '0 1px 14px #10b6ef99'
    }}>NetReconUltra</h1>
    <div style={{ flex: 1 }} />
    <nav style={{ fontSize: 18 }}>
      <a href="#dashboard" style={{ marginRight: 22 }}>Dashboard</a>
      <a href="#devices" style={{ marginRight: 22 }}>Geräte</a>
      <a href="#about">Über</a>
    </nav>
    <ThemeSwitch />
  </header>
);
