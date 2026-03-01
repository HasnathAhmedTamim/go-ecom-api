import { useMutation } from '@tanstack/react-query'
import api from '../api'

export function useCreateOrder() {
  return useMutation({
    mutationFn: async ({ items, address }) => {
      // Map items array to products map { id: qty }
      const products = {}
      for (const it of items) {
        products[it.id] = (products[it.id] || 0) + (it.qty || 1)
      }
      const payload = { products }
      const { data } = await api.post('/api/user/orders', payload)
      return data
    },
  })
}

export default useCreateOrder
