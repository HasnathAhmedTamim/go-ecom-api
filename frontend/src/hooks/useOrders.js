import { useQuery } from '@tanstack/react-query'
import api from '../api'

export function useOrders() {
  return useQuery({
    queryKey: ['orders'],
    queryFn: async () => {
      const { data } = await api.get('/api/user/orders')
      return data
    },
  })
}

export default useOrders
