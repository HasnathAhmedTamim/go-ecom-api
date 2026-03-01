# Frontend (Landing demo)

This folder contains the frontend for the Micro‑SaaS Dashboard landing/demo built with Vite, React and Tailwind.

Quick start

1. Install dependencies:

```bash
cd frontend
npm install
```

2. Run the dev server:

```bash
npm run dev
```

- The Vite dev server runs on `http://localhost:5173` (or next available port).
- The marketing landing is available at `/landing`.

Mock API (dev-only)

- A lightweight dev middleware (`src/mocks/devApiMiddleware.js`) returns mock JSON for `/api/*` endpoints when the backend is unavailable. This lets you develop and preview the landing/demo without running the Go backend.

Backend notes

- The backend (Go server) uses `go-sqlite3`, which requires CGO. On Windows you must have a C toolchain available and run with CGO enabled. Example (PowerShell):

```powershell
setx CGO_ENABLED 1
# ensure a C toolchain (MSYS2 or MinGW) is installed and in PATH
cd ../backend
go run cmd/server/main.go
```

If you prefer not to run the backend locally, the frontend mock covers basic product endpoints used by the landing and demo pages.

Packaging / ThemeForest notes

- Include the following in the final package: compiled/minified `dist` build, `index.html` demo, `src` (components + styles), `README.md`, and Figma sources (if available).
- Provide clear install and demo instructions, list features (dark mode, RTL, responsive components, Storybook), and include screenshots and an accessible demo URL.

Files of interest

- `src/pages/Landing.jsx` — landing page entry
- `src/components/LandingHero.jsx` — hero section
- `src/components/Features.jsx` — features grid
- `src/mocks/devApiMiddleware.js` — dev mock middleware
- `vite.config.js` — Vite config with proxy + dev mock

Screenshots

- `screenshots/hero-1.svg` — hero / dashboard preview mock
- `screenshots/feature-grid.svg` — feature grid preview

# React + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) (or [oxc](https://oxc.rs) when used in [rolldown-vite](https://vite.dev/guide/rolldown)) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## React Compiler

The React Compiler is not enabled on this template because of its impact on dev & build performances. To add it, see [this documentation](https://react.dev/learn/react-compiler/installation).

## Expanding the ESLint configuration

If you are developing a production application, we recommend using TypeScript with type-aware lint rules enabled. Check out the [TS template](https://github.com/vitejs/vite/tree/main/packages/create-vite/template-react-ts) for information on how to integrate TypeScript and [`typescript-eslint`](https://typescript-eslint.io) in your project.
