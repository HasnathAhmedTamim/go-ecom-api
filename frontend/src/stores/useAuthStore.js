import { create } from 'zustand'

const initialToken = typeof window !== 'undefined' ? localStorage.getItem('auth_token') : null
const initialUser = typeof window !== 'undefined' ? JSON.parse(localStorage.getItem('auth_user') || 'null') : null

export const useAuthStore = create((set) => ({
  token: initialToken,
  user: initialUser,
  setAuth: (user, token) => {
    localStorage.setItem('auth_token', token)
    localStorage.setItem('auth_user', JSON.stringify(user))
    set({ user, token })
  },
  clearAuth: () => {
    localStorage.removeItem('auth_token')
    localStorage.removeItem('auth_user')
    set({ user: null, token: null })
  },
}))

export default useAuthStore
