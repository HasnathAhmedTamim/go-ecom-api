import { useCartStore } from '../stores/useCartStore'

export default function Cart() {
  const items = useCartStore((s) => s.items)
  const remove = useCartStore((s) => s.remove)

  if (!items.length) return <div>Your cart is empty</div>

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4">Cart</h1>
      <div className="bg-white rounded-lg shadow p-4">
        <ul>
          {items.map((it) => (
            <li key={it.id} className="flex items-center justify-between py-3 border-b last:border-b-0">
              <div className="flex items-center gap-4">
                <div className="w-16 h-12 bg-gray-100 rounded flex items-center justify-center">Img</div>
                <div>
                  <div className="font-medium">{it.name}</div>
                  <div className="text-sm text-gray-600">${it.price}</div>
                </div>
              </div>
              <div className="flex items-center gap-3">
                <button className="text-red-600" onClick={() => remove(it.id)}>Remove</button>
              </div>
            </li>
          ))}
        </ul>
        <div className="mt-4 text-right font-semibold">Total: ${items.reduce((s, i) => s + (i.price || 0), 0)}</div>
      </div>
    </section>
  )
}
