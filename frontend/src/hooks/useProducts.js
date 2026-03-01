import { useQuery } from '@tanstack/react-query'
import api from '../api'

export function useProducts() {
  return useQuery({
    queryKey: ['products'],
    queryFn: async () => {
      const { data } = await api.get('/api/products')
      // backend returns { items, page, limit, total }
      // return items array for components expecting an array
      return data?.items ?? data
    },
  })
}
