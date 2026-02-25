import axios from 'axios'

// Use relative base URL so requests go to the Vite dev server and
// are proxied to the backend at /api during development.
const api = axios.create({
  baseURL: '',
  withCredentials: true,
})

export default api
