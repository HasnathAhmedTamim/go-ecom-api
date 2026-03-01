import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
        onError: (err, req, res) => {
          if (res && !res.writableEnded) {
            res.writeHead && res.writeHead(502, { 'Content-Type': 'text/plain' })
            res.end('Backend unavailable (dev proxy)')
          }
        },
      },
    },
    // Enable the mock middleware when running Vite in dev mode
    middlewareMode: false,
  },
})
