import { useLocation, Navigate } from 'react-router-dom'
import { useAuthStore } from '../stores/useAuthStore'

export default function AdminRoute({ children }) {
  const user = useAuthStore((s) => s.user)
  const location = useLocation()
  const role = user?.role
  if (!user) return <Navigate to="/login" state={{ from: location }} replace />
  if (role !== 'admin') return <Navigate to="/" replace />
  return children
}
