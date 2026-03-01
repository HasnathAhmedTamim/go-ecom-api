import { useLocation, Navigate } from 'react-router-dom'
import { useAuthStore } from '../stores/useAuthStore'

export default function PrivateRoute({ children }) {
  const token = useAuthStore((s) => s.token)
  const location = useLocation()
  if (!token) {
    return <Navigate to="/login" state={{ from: location }} replace />
  }
  return children
}
