import { useQuery } from '@tanstack/react-query'
import api from '../api'

export function useProducts(filters = {}) {
  return useQuery({
    queryKey: ['products', filters],
    queryFn: async () => {
      const params = new URLSearchParams()
      if (filters.q) params.set('q', filters.q)
      if (filters.minPrice) params.set('min_price', filters.minPrice)
      if (filters.maxPrice) params.set('max_price', filters.maxPrice)
      if (filters.category) params.set('category', filters.category)
      const url = '/api/products' + (params.toString() ? `?${params.toString()}` : '')
      const { data } = await api.get(url)
      return data?.items ?? data
    },
  })
}
