import { useParams } from 'react-router-dom'
import { useEffect, useState } from 'react'
import api from '../api'
import { useCartStore } from '../stores/useCartStore'
import Button from '../components/Button'
import useToastStore from '../stores/useToastStore'

export default function Product() {
  const { id } = useParams()
  const [product, setProduct] = useState(null)
  const [loading, setLoading] = useState(true)
  const [err, setErr] = useState(null)
  const add = useCartStore((s) => s.add)
  const push = useToastStore((s) => s.push)

  useEffect(() => {
    let mounted = true
    api
      .get(`/api/products/${id}`)
      .then((res) => mounted && setProduct(res.data))
      .catch((e) => mounted && setErr(e.message))
      .finally(() => mounted && setLoading(false))
    return () => (mounted = false)
  }, [id])

  if (loading) return <div className="text-white">Loading...</div>
  if (err) return <div className="text-red-400">Error: {err}</div>
  if (!product) return <div className="text-gray-300">Not found</div>

  return (
    <article className="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div className="md:col-span-2 bg-black/60 rounded-lg shadow-neon p-6">
        <div className="h-64 bg-gradient-to-br from-white/3 to-white/2 rounded flex items-center justify-center mb-4">
          <img src={(product.image || '/screenshots/placeholder.svg').replace('.jpg', '.svg')} alt={product.name} className="h-56 object-contain" />
        </div>
        <h1 className="text-2xl font-semibold text-white">{product.name}</h1>
        <p className="mt-3 text-gray-300">{product.description}</p>
      </div>
      <aside className="md:col-span-1 bg-black/60 rounded-lg shadow-neon p-6 h-fit border border-white/5">
        <div className="text-xl font-semibold text-neon-cyan">${product.price}</div>
        <div className="mt-4">
          <Button onClick={() => { add({ id: product.id, name: product.name, price: product.price }); push({ type: 'success', title: 'Added to cart' }) }}>
            Add to cart
          </Button>
        </div>
      </aside>
    </article>
  )
}
