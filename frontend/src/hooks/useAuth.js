import { useEffect } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import { useAuthStore } from '../stores/useAuthStore'

function parseJwt(token) {
  try {
    const parts = token.split('.')
    if (parts.length < 2) return null
    const payload = parts[1]
    const json = atob(payload.replace(/-/g, '+').replace(/_/g, '/'))
    return JSON.parse(decodeURIComponent(escape(json)))
  } catch {
    return null
  }
}

export function useAuth() {
  const token = useAuthStore((s) => s.token)
  const user = useAuthStore((s) => s.user)
  const setAuth = useAuthStore((s) => s.setAuth)
  const clearAuth = useAuthStore((s) => s.clearAuth)
  const qc = useQueryClient()

  useEffect(() => {
    if (!token) return
    // Prefer refreshing user from server if token present
    let mounted = true
    if (!user) {
      ;(async () => {
        try {
          const res = await fetch('/api/auth/me', {
            headers: { Authorization: `Bearer ${token}` },
          })
          if (!mounted) return
          if (res.ok) {
            const json = await res.json()
            const srvUser = json.user
            if (srvUser) setAuth(srvUser, token)
            return
          }
          // fallback to decode claims
          const claims = parseJwt(token)
          if (claims) {
            const minimal = { id: claims.user_id || claims.sub || null, role: claims.role || null }
            setAuth(minimal, token)
          }
        } catch {
          const claims = parseJwt(token)
          if (claims) {
            const minimal = { id: claims.user_id || claims.sub || null, role: claims.role || null }
            setAuth(minimal, token)
          }
        }
      })()
    }
    return () => {
      mounted = false
    }
  }, [token, setAuth, user])

  function logout() {
    clearAuth()
    // clear react-query cache
    try {
      qc.clear()
    } catch {
      /* ignore */
    }
  }

  return { token, user, logout }
}

export default useAuth
