import { useState } from 'react'
import { useNavigate, useLocation, Link } from 'react-router-dom'
import api from '../api'
import { useAuthStore } from '../stores/useAuthStore'
import Button from '../components/Button'
import useToastStore from '../stores/useToastStore'

export default function Login() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState(null)
  const navigate = useNavigate()
  const location = useLocation()
  const setAuth = useAuthStore((s) => s.setAuth)

  const from = location.state?.from?.pathname || '/'

  const submit = async (e) => {
    e.preventDefault()
    setError(null)
    try {
      const res = await api.post('/api/auth/login', { email, password })
      const { user, token } = res.data
      setAuth(user, token)
      useToastStore.getState().push({ type: 'success', title: 'Signed in' })
      navigate(from, { replace: true })
    } catch (err) {
      setError(err?.response?.data?.error || err.message || 'Login failed')
      useToastStore.getState().push({ type: 'error', title: 'Login failed', message: setError })
    }
  }

  return (
    <section className="max-w-md mx-auto mt-8 bg-white rounded-lg shadow p-6">
      <h1 className="text-2xl font-semibold mb-4">Sign in to your account</h1>
      <form onSubmit={submit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium mb-1">Email</label>
          <input className="border p-2 w-full rounded" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="you@example.com" />
        </div>
        <div>
          <label className="block text-sm font-medium mb-1">Password</label>
          <input type="password" className="border p-2 w-full rounded" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Your password" />
        </div>

        <div className="flex items-center justify-between">
          <Button type="submit">Login</Button>
          <Link to="/register" className="text-sm text-blue-600">Create account</Link>
        </div>

        {error && <div className="text-red-600 mt-2">{error}</div>}
      </form>
    </section>
  )
}
