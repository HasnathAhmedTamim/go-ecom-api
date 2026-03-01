import { Link } from 'react-router-dom'
import { useProducts } from '../hooks/useProducts'
import { useCartStore } from '../stores/useCartStore'
import useToastStore from '../stores/useToastStore'
import Button from '../components/Button'

export default function Products() {
  const { data, isLoading, error } = useProducts()
  const add = useCartStore((s) => s.add)
  const push = useToastStore((s) => s.push)

  if (isLoading) return <div>Loading products...</div>
  if (error) return <div>Error loading products</div>

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-6">Products</h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {data?.map((p) => (
          <div key={p.id} className="bg-white rounded-lg shadow p-4 flex flex-col">
            <div className="h-44 bg-gradient-to-br from-gray-50 to-gray-100 rounded mb-4 flex items-center justify-center text-gray-400">
              <img src={p.image || '/screenshots/hero-1.svg'} alt={p.name} className="h-36 object-contain" />
            </div>
            <h2 className="font-medium text-lg">{p.name}</h2>
            <p className="text-sm text-gray-600 flex-1 mt-2 line-clamp-3">{p.description || 'High-quality components and integrations for product teams.'}</p>
            <div className="mt-4 flex items-center justify-between">
              <div className="font-semibold">${p.price ?? '—'}</div>
              <div className="flex gap-2 items-center">
                <Link to={`/products/${p.id}`} className="text-indigo-600 hover:underline">View</Link>
                <Button onClick={() => { add({ id: p.id, name: p.name, price: p.price }); push({ type: 'success', title: 'Added to cart' }) }}>Add</Button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </section>
  )
}
