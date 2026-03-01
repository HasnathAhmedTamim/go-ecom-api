import { useEffect, useState } from 'react'
import api from '../api'

export default function AdminOrders() {
  const [orders, setOrders] = useState([])
  const [loading, setLoading] = useState(true)
  const [err, setErr] = useState(null)

  useEffect(() => {
    let mounted = true
    api.get('/api/admin/orders').then((res) => mounted && setOrders(res.data)).catch((e) => mounted && setErr(e.message)).finally(() => mounted && setLoading(false))
    return () => (mounted = false)
  }, [])

  if (loading) return <div>Loading...</div>
  if (err) return <div>Error: {err}</div>

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4">Manage Orders</h1>
      <ul>
        {orders.map((o) => (
          <li key={o.id} className="py-2 border-b flex justify-between items-center">
            <div>
              <div className="font-medium">Order {o.id} - {o.status}</div>
              <div className="text-sm text-gray-600">User: {o.user_id}</div>
            </div>
            <div className="flex gap-2">
              <button className="text-indigo-600">View</button>
              <button className="text-green-600">Mark shipped</button>
            </div>
          </li>
        ))}
      </ul>
    </section>
  )
}
