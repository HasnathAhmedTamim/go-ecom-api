import { useEffect, useState } from 'react'
import api from '../api'
import useToastStore from '../stores/useToastStore'
import { useAuthStore } from '../stores/useAuthStore'

export default function AdminProducts() {
  const [products, setProducts] = useState([])
  const [loading, setLoading] = useState(true)
  const [err, setErr] = useState(null)
  const [creating, setCreating] = useState(false)
  const [name, setName] = useState('')
  const [price, setPrice] = useState('')
  const [stock, setStock] = useState('')

  useEffect(() => {
    let mounted = true
    api
      .get('/api/products')
      .then((res) => {
        const data = res.data
        const items = data && data.items ? data.items : data
        if (mounted) setProducts(items)
      })
      .catch((err) => mounted && setErr(err.message))
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
    } catch {
      push({ type: 'error', title: 'Delete failed' })
    }
  }

  const createProduct = async (e) => {
    e.preventDefault()
    if (!name || !price) return push({ type: 'error', title: 'Name and price required' })
    try {
      setCreating(true)
      const payload = { name, price: Number(price), stock: Number(stock || 0) }
      const res = await api.post('/api/admin/products', payload)
      const created = res.data
      setProducts((p) => [created, ...p])
      window.dispatchEvent(new Event('products:changed'))
      push({ type: 'success', title: 'Product created' })
      setName('')
      setPrice('')
      setStock('')
    } catch {
      push({ type: 'error', title: 'Create failed' })
    } finally {
      setCreating(false)
    }
  }

  if (loading) return <div>Loading...</div>
  if (err) return <div>Error: {err}</div>

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4">Manage Products</h1>
      {user?.role === 'admin' && (
        <form onSubmit={createProduct} className="mb-4 p-4 border rounded bg-white">
          <div className="grid grid-cols-3 gap-2">
            <input className="border p-2" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
            <input className="border p-2" placeholder="Price" value={price} onChange={(e) => setPrice(e.target.value)} />
            <input className="border p-2" placeholder="Stock" value={stock} onChange={(e) => setStock(e.target.value)} />
          </div>
          <div className="mt-2">
            <button className="px-3 py-1 bg-green-600 text-white rounded" disabled={creating}>{creating ? 'Creating...' : 'Create Product'}</button>
          </div>
        </form>
      )}
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
