import { useMutation } from '@tanstack/react-query'
import api from '../api'

export function useLogin() {
  return useMutation({
    mutationFn: async (credentials) => {
      const { data } = await api.post('/api/auth/login', credentials)
      return data
    },
  })
}

export default useLogin
