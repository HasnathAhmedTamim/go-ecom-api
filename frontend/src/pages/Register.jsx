import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import api from '../api'
import Button from '../components/Button'
import useToastStore from '../stores/useToastStore'

export default function Register() {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState(null)
  const navigate = useNavigate()

  const submit = async (e) => {
    e.preventDefault()
    setError(null)
    try {
      await api.post('/api/auth/register', { name, email, password })
      useToastStore.getState().push({ type: 'success', title: 'Account created' })
      navigate('/login')
    } catch (err) {
      setError(err?.response?.data?.error || err.message || 'Registration failed')
      useToastStore.getState().push({ type: 'error', title: 'Registration failed' })
    }
  }

  return (
    <section className="max-w-md mx-auto mt-8 bg-white rounded-lg shadow p-6">
      <h1 className="text-2xl font-semibold mb-4">Create an account</h1>
      <form onSubmit={submit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium mb-1">Name</label>
          <input className="border p-2 w-full rounded" value={name} onChange={(e) => setName(e.target.value)} placeholder="Your name" />
        </div>
        <div>
          <label className="block text-sm font-medium mb-1">Email</label>
          <input className="border p-2 w-full rounded" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="you@example.com" />
        </div>
        <div>
          <label className="block text-sm font-medium mb-1">Password</label>
          <input type="password" className="border p-2 w-full rounded" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Choose a password" />
        </div>
        <div className="flex items-center justify-between">
          <Button type="submit">Create account</Button>
          <Link to="/login" className="text-sm text-blue-600">Already have an account?</Link>
        </div>
        {error && <div className="text-red-600 mt-2">{error}</div>}
      </form>
    </section>
  )
}
