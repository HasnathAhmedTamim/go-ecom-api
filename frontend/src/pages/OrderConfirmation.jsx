import { useEffect, useState } from 'react'
import { useParams, Link } from 'react-router-dom'
import api from '../api'

export default function OrderConfirmation() {
  const { id } = useParams()
  const [order, setOrder] = useState(null)
  const [loading, setLoading] = useState(true)
  const [err, setErr] = useState(null)

  useEffect(() => {
    let mounted = true
    api
      .get(`/api/user/orders/${id}`)
      .then((res) => mounted && setOrder(res.data))
      .catch((e) => mounted && setErr(e.message))
      .finally(() => mounted && setLoading(false))
    return () => (mounted = false)
  }, [id])

  if (loading) return <div>Loading...</div>
  if (err) return <div>Error: {err}</div>
  if (!order) return <div>Order not found</div>

  return (
    <section className="max-w-xl mx-auto mt-8 bg-white rounded-lg shadow p-6">
      <h1 className="text-2xl font-semibold mb-4">Order Placed</h1>
      <div className="mb-2">Order ID: <strong>{order.id}</strong></div>
      <div className="mb-2">Status: {order.status}</div>
      <div className="mb-4">Products:</div>
      <ul className="mb-4">
        {Object.entries(order.products || {}).map(([pid, qty]) => (
          <li key={pid} className="py-1">{pid} — qty: {qty}</li>
        ))}
      </ul>
      <Link to="/orders" className="text-blue-600">View all orders</Link>
    </section>
  )
}
