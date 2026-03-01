import { Link } from 'react-router-dom'
import useAuth from '../hooks/useAuth'
import { useAuthStore } from '../stores/useAuthStore'
import { useCartStore } from '../stores/useCartStore'
import useToastStore from '../stores/useToastStore'

export default function Header() {
  const { logout } = useAuth()
  const storeUser = useAuthStore((s) => s.user)
  const cartCount = useCartStore((s) => s.items.length)
  const displayName = storeUser ? (storeUser.name || storeUser.email || '') : ''
  const displayNameCap = displayName ? displayName.charAt(0).toUpperCase() + displayName.slice(1) : ''
  const initial = displayName ? displayName.charAt(0).toUpperCase() : 'U'

  const isAdmin = storeUser?.role === 'admin'

  return (
    <header className="bg-white shadow">
      <div className="container mx-auto px-4 py-4 flex items-center justify-between">
        <Link to="/" className="flex items-center gap-3">
          <div className="w-10 h-10 bg-gradient-to-br from-indigo-500 to-blue-500 rounded-md flex items-center justify-center text-white font-bold">MS</div>
          <div>
            <div className="text-lg font-semibold">MicroSaaS</div>
            <div className="text-xs text-gray-500">Dashboard demo</div>
          </div>
        </Link>

        <nav className="hidden md:flex items-center gap-6">
          <Link to="/products" className="text-gray-700 hover:text-gray-900">Products</Link>
          <Link to="/cart" className="text-gray-700 hover:text-gray-900">Cart ({cartCount})</Link>
          <Link to="/orders" className="text-gray-700 hover:text-gray-900">Orders</Link>
          {isAdmin && <Link to="/admin" className="text-indigo-600 font-medium">Admin</Link>}

          {storeUser ? (
            <div className="flex items-center gap-4">
              <div className="flex items-center gap-2">
                <div className="w-8 h-8 bg-gray-200 rounded-full flex items-center justify-center text-sm text-gray-700">{initial}</div>
                <div className="text-sm text-gray-700">{displayNameCap}</div>
              </div>
              <button
                onClick={() => {
                  if (confirm('Logout?')) {
                    logout()
                    useToastStore.getState().push({ type: 'success', title: 'Signed out' })
                  }
                }}
                className="text-sm text-red-600"
                aria-label="Sign out"
              >
                Logout
              </button>
            </div>
          ) : (
            <div className="flex items-center gap-4">
              <Link to="/login" className="text-gray-700 hover:text-gray-900">Login</Link>
              <Link to="/register" className="text-white bg-blue-600 px-3 py-1 rounded-md">Sign up</Link>
            </div>
          )}
        </nav>

        {/* Mobile menu placeholder */}
        <div className="md:hidden">
          <Link to="/cart" className="text-gray-700">Cart ({cartCount})</Link>
        </div>
      </div>
    </header>
  )
}
