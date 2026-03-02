import { Link } from 'react-router-dom'
import { useProducts } from '../hooks/useProducts'
import { useCartStore } from '../stores/useCartStore'
import useToastStore from '../stores/useToastStore'
import Button from '../components/Button'
import { useState } from 'react'

export default function Products() {
  const [category, setCategory] = useState('')
  const { data, isLoading, error } = useProducts({ category })
  const add = useCartStore((s) => s.add)
  const push = useToastStore((s) => s.push)

  if (isLoading) return <div>Loading products...</div>
  if (error) return <div>Error loading products</div>

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-6 text-white">Products</h1>
        <div className="mb-4 flex gap-2 items-center">
          <label className="text-sm text-gray-300">Category:</label>
          <select value={category} onChange={(e) => setCategory(e.target.value)} className="border p-2 rounded bg-black text-gray-300 border-white/10">
            <option value="">All</option>
            <option value="Audio">Audio</option>
            <option value="Keyboards">Keyboards</option>
            <option value="Mice">Mice</option>
            <option value="Chairs">Chairs</option>
            <option value="Controllers">Controllers</option>
          </select>
        </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {data?.map((p) => (
          <div key={p.id} className="bg-black/60 rounded-lg shadow-neon p-4 flex flex-col border border-white/5">
            <div className="h-44 bg-gradient-to-br from-white/3 to-white/2 rounded mb-4 flex items-center justify-center text-gray-400">
              <img src={p.image || '/screenshots/hero-1.svg'} alt={p.name} className="h-36 object-contain" />
            </div>
            <h2 className="font-medium text-lg text-white">{p.name}</h2>
            <p className="text-sm text-gray-300 flex-1 mt-2 line-clamp-3">{p.description || 'High-quality gaming gear and accessories.'}</p>
            <div className="mt-4 flex items-center justify-between">
              <div className="font-semibold text-neon-cyan">${p.price ?? '—'}</div>
              <div className="flex gap-2 items-center">
                <Link to={`/products/${p.id}`} className="text-neon-cyan hover:underline">View</Link>
                <Button onClick={() => { add({ id: p.id, name: p.name, price: p.price }); push({ type: 'success', title: 'Added to cart' }) }}>Add</Button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </section>
  )
}
