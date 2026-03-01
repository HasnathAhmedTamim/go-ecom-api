import { useEffect, useState } from 'react'
import api from '../api'
import useToastStore from '../stores/useToastStore'
import { useAuthStore } from '../stores/useAuthStore'

export default function AdminProducts() {
  const [products, setProducts] = useState([])
  const [loading, setLoading] = useState(true)
  const [err, setErr] = useState(null)

  useEffect(() => {
    let mounted = true
    api
      .get('/api/products')
      .then((res) => {
        const data = res.data
        const items = data && data.items ? data.items : data
        if (mounted) setProducts(items)
      })
      .catch((e) => mounted && setErr(e.message))
      .finally(() => mounted && setLoading(false))
    return () => (mounted = false)
  }, [])

  const push = useToastStore((s) => s.push)
  const user = useAuthStore((s) => s.user)

  const del = async (id) => {
    if (!confirm('Delete product?')) return
    try {
      await api.delete(`/api/admin/products/${id}`)
      setProducts((p) => p.filter((x) => x.id !== id))
      window.dispatchEvent(new Event('products:changed'))
      push({ type: 'success', title: 'Product deleted' })
    } catch (e) {
      push({ type: 'error', title: 'Delete failed' })
    }
  }

  if (loading) return <div>Loading...</div>
  if (err) return <div>Error: {err}</div>

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4">Manage Products</h1>
      <ul>
        {products.map((p) => (
          <li key={p.id} className="flex justify-between py-2">
            <div>
              <div className="font-medium">{p.name}</div>
              <div className="text-sm text-gray-600">${p.price}</div>
            </div>
            <div>
              {user?.role === 'admin' ? (
                <>
                  <button className="text-blue-600 mr-2">Edit</button>
                  <button className="text-red-600" onClick={() => del(p.id)}>Delete</button>
                </>
              ) : (
                <span className="text-sm text-gray-500">Admin only</span>
              )}
            </div>
          </li>
        ))}
      </ul>
    </section>
  )
}
