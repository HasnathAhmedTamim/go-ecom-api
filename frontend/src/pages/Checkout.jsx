import { useMemo, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useCartStore } from '../stores/useCartStore'
import useCreateOrder from '../hooks/useCreateOrder'
import useCheckout from '../hooks/useCheckout'
import Button from '../components/Button'
import useToastStore from '../stores/useToastStore'

export default function Checkout() {
  const items = useCartStore((s) => s.items)
  const clear = useCartStore((s) => s.clear)
  const navigate = useNavigate()
  const [address, setAddress] = useState('')
  const createOrder = useCreateOrder()
  const checkout = useCheckout()

  const total = useMemo(() => items.reduce((s, i) => s + (i.price || 0), 0), [items])

  const submit = async (e) => {
    e?.preventDefault()
    if (!items.length) return alert('Cart is empty')
    try {
      // Use the mock checkout flow which returns a payment_url and order
      const resp = await checkout.mutateAsync({ items, address })
      clear()
      useToastStore.getState().push({ type: 'success', title: 'Order placed' })
      // Redirect user to mock payment URL in a new tab
      if (resp.payment_url) {
        window.open(resp.payment_url, '_blank')
      }
      navigate(`/order/${resp.order.id}`)
    } catch {
      useToastStore.getState().push({ type: 'error', title: 'Order failed' })
    }
  }

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4 text-white">Checkout</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="md:col-span-2 bg-black/60 rounded-lg shadow-neon p-4 border border-white/5 text-gray-300">
          <div className="font-medium text-white">Items:</div>
          <ul>
            {items.map((it) => (
              <li key={it.id} className="py-1 border-b last:border-b-0 border-white/5">{it.name} — <span className="text-neon-cyan">${it.price}</span></li>
            ))}
          </ul>
          <div className="mt-2 text-white">Total: <span className="text-neon-cyan">${total}</span></div>
        </div>

        <aside className="md:col-span-1 bg-black/60 rounded-lg shadow-neon p-4 border border-white/5">
          <form onSubmit={submit} className="flex flex-col">
            <label className="block mb-2 text-gray-300">Shipping address</label>
            <textarea required className="border p-2 w-full mb-4 rounded bg-black text-gray-200 border-white/10" value={address} onChange={(e) => setAddress(e.target.value)} />
            <Button type="submit">Place order — ${total}</Button>
          </form>
        </aside>
      </div>
    </section>
  )
}
