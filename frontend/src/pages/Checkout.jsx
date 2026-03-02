import { useMemo, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useCartStore } from '../stores/useCartStore'
import useCreateOrder from '../hooks/useCreateOrder'
import Button from '../components/Button'
import useToastStore from '../stores/useToastStore'

export default function Checkout() {
  const items = useCartStore((s) => s.items)
  const clear = useCartStore((s) => s.clear)
  const navigate = useNavigate()
  const [address, setAddress] = useState('')
  const createOrder = useCreateOrder()

  const total = useMemo(() => items.reduce((s, i) => s + (i.price || 0), 0), [items])

  const submit = async (e) => {
    e?.preventDefault()
    if (!items.length) return alert('Cart is empty')
    try {
      const created = await createOrder.mutateAsync({ items, address })
      clear()
      useToastStore.getState().push({ type: 'success', title: 'Order placed' })
      navigate(`/order/${created.id}`)
    } catch {
      useToastStore.getState().push({ type: 'error', title: 'Order failed' })
    }
  }

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4">Checkout</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="md:col-span-2 bg-white rounded-lg shadow p-4">
          <div className="font-medium">Items:</div>
          <ul>
            {items.map((it) => (
              <li key={it.id} className="py-1 border-b last:border-b-0">{it.name} — ${it.price}</li>
            ))}
          </ul>
          <div className="mt-2">Total: ${total}</div>
        </div>

        <aside className="md:col-span-1 bg-white rounded-lg shadow p-4">
          <form onSubmit={submit} className="flex flex-col">
            <label className="block mb-2">Shipping address</label>
            <textarea required className="border p-2 w-full mb-4 rounded" value={address} onChange={(e) => setAddress(e.target.value)} />
            <Button type="submit">Place order — ${total}</Button>
          </form>
        </aside>
      </div>
    </section>
  )
}
