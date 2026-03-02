import axios from 'axios'

// Allow overriding API base URL via Vite env var `VITE_API_BASE`.
// Default to empty string so dev proxy/mocks work.
const base = import.meta.env.VITE_API_BASE ?? ''
const api = axios.create({ baseURL: base, withCredentials: true })

// Attach Authorization header from localStorage token when present
api.interceptors.request.use((cfg) => {
  try {
    const token = localStorage.getItem('auth_token')
    if (token) cfg.headers = { ...(cfg.headers || {}), Authorization: `Bearer ${token}` }
  } catch {
    // ignore
  }
  return cfg
})

export default api
