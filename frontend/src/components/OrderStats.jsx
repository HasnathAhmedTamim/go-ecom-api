import { useEffect, useState, useCallback } from 'react'
import api from '../api'

export default function OrderStats() {
  const [loading, setLoading] = useState(true)
  const [stats, setStats] = useState({ totalRevenue: 0, count: 0, byStatus: {} })

  const fetchStats = useCallback(async () => {
    setLoading(true)
    try {
      const { data } = await api.get('/api/admin/orders')
      const orders = Array.isArray(data) ? data : []
      const totalRevenue = orders.reduce((s, o) => s + (Number(o.total) || 0), 0)
      const byStatus = orders.reduce((acc, o) => {
        const st = o.status || 'unknown'
        acc[st] = (acc[st] || 0) + 1
        return acc
      }, {})
      setStats({ totalRevenue, count: orders.length, byStatus })
    } catch (e) {
      console.error('OrderStats fetch error', e)
      setStats({ totalRevenue: 0, count: 0, byStatus: {} })
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    fetchStats()
    const handler = () => fetchStats()
    window.addEventListener('products:changed', handler)
    return () => window.removeEventListener('products:changed', handler)
  }, [fetchStats])

  if (loading) return <div className="text-gray-300">Loading stats...</div>

  return (
    <div className="p-4 border rounded bg-black/60 text-gray-300 border-white/5">
      <h2 className="text-lg font-medium text-white">Order Statistics</h2>
      <div className="mt-2 text-sm">
        <p><strong>Total revenue:</strong> <span className="text-neon-cyan">${stats.totalRevenue.toFixed(2)}</span></p>
        <p><strong>Orders:</strong> {stats.count}</p>
        <div className="mt-2">
          <strong>By status:</strong>
          <ul className="list-disc ml-6">
            {Object.entries(stats.byStatus).map(([k, v]) => (
              <li key={k}>{k}: {v}</li>
            ))}
          </ul>
        </div>
        <div className="mt-3">
          <button className="px-3 py-1 bg-neon-pink text-black rounded" onClick={fetchStats}>Refresh</button>
        </div>
      </div>
    </div>
  )
}
