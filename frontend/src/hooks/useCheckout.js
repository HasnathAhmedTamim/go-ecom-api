import { useMutation } from '@tanstack/react-query'
import api from '../api'

export function useCheckout() {
  return useMutation({
    mutationFn: async ({ items, address }) => {
      const payload = { items, address }
      const { data } = await api.post('/api/user/checkout', payload)
      return data
    },
  })
}

export default useCheckout
