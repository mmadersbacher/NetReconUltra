import React from 'react';
import { Header } from '../components/Header';
import { Footer } from '../components/Footer';

export const MainLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => (
  <div style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
    <Header />
    <main style={{ flex: 1, width: '100%', maxWidth: 1200, margin: '0 auto' }}>
      {children}
    </main>
    <Footer />
  </div>
);
