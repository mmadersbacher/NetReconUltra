import React from 'react';

export const Footer: React.FC = () => (
  <footer style={{
    textAlign: 'center', padding: '2rem 0 1rem 0', color: 'var(--color-muted)', fontSize: 15
  }}>
    <hr style={{ margin: '2rem 0 1rem 0', borderColor: '#223' }} />
    <div>NetReconUltra © 2025 – <a href="https://github.com/mmadersbacher/NetReconUltra" target="_blank">GitHub Repo</a></div>
    <div style={{ marginTop: 6 }}>Created by Mario Madersbacher</div>
  </footer>
);
