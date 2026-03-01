import { useQuery } from '@tanstack/react-query'
import api from '../api'

export function useProduct(id) {
  return useQuery({
    queryKey: ['product', id],
    queryFn: async () => {
      const { data } = await api.get(`/api/products/${id}`)
      return data
    },
    enabled: !!id,
  })
}

export default useProduct
