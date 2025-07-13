import React, { useRef, useEffect } from "react";

export const TechNetBackground: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    let ctx = canvas.getContext('2d');
    if (!ctx) return;

    let width = canvas.width = canvas.offsetWidth;
    let height = canvas.height = canvas.offsetHeight;
    const nodes = Array.from({ length: 22 }, () => ({
      x: Math.random() * width, y: Math.random() * height,
      vx: (Math.random() - 0.5) * 0.5, vy: (Math.random() - 0.5) * 0.5
    }));

    function draw() {
      // SAFETY: Context immer checken
      if (!ctx) return;
      ctx.clearRect(0, 0, width, height);

      // Draw links
      for (let i = 0; i < nodes.length; ++i) {
        for (let j = i + 1; j < nodes.length; ++j) {
          let a = nodes[i], b = nodes[j];
          let dist = Math.hypot(a.x - b.x, a.y - b.y);
          if (dist < 140 && ctx) {
            ctx.strokeStyle = `rgba(16,182,239,${0.14 + 0.13 * (1 - dist / 140)})`;
            ctx.lineWidth = 1.6 - dist / 210;
            ctx.beginPath(); ctx.moveTo(a.x, a.y); ctx.lineTo(b.x, b.y); ctx.stroke();
          }
        }
      }
      // Draw nodes
      for (let n of nodes) {
        if (!ctx) continue;
        ctx.beginPath();
        ctx.arc(n.x, n.y, 4.6, 0, 2 * Math.PI);
        ctx.fillStyle = "#10b6ef";
        ctx.shadowColor = "#00e9ff";
        ctx.shadowBlur = 9;
        ctx.fill();
        ctx.shadowBlur = 0;
      }
    }

    function animate() {
      for (let n of nodes) {
        n.x += n.vx;
        n.y += n.vy;
        if (n.x < 0 || n.x > width) n.vx *= -1;
        if (n.y < 0 || n.y > height) n.vy *= -1;
      }
      draw();
      requestAnimationFrame(animate);
    }

    animate();

    window.addEventListener('resize', () => {
      width = canvas.width = canvas.offsetWidth;
      height = canvas.height = canvas.offsetHeight;
      ctx = canvas.getContext('2d'); // Im Falle von Resize Kontext neu holen
    });

    // Kein Cleanup nötig für AnimationFrame, da Seite nie entlädt
    // Wenn du willst, kannst du ein CancelToken für requestAnimationFrame bauen

  }, []);

  return (
    <canvas
      ref={canvasRef}
      width={820}
      height={320}
      style={{
        position: 'absolute', left: 0, top: 0, width: '100%', height: '100%',
        zIndex: 0, borderRadius: 26, opacity: 0.55,
        pointerEvents: 'none', filter: 'blur(1px)'
      }}
      aria-hidden="true"
    />
  );
};
