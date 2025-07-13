import React from 'react';
import { Sun, Moon } from 'lucide-react';

export const ThemeSwitch: React.FC = () => {
  const [theme, setTheme] = React.useState<'dark' | 'light'>(
    (localStorage.getItem('theme') as 'dark' | 'light') || 'dark'
  );

  React.useEffect(() => {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
  }, [theme]);

  return (
    <button
      aria-label="Theme wechseln"
      onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
      style={{
        background: 'none',
        border: 'none',
        marginLeft: 14,
        cursor: 'pointer',
        color: 'var(--color-accent)',
        fontSize: 22,
        transition: 'color .18s',
        verticalAlign: 'middle'
      }}
    >
      {theme === 'dark' ? <Sun size={21} /> : <Moon size={21} />}
    </button>
  );
};
