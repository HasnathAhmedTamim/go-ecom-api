import { useCartStore } from '../stores/useCartStore'

export default function Cart() {
  const items = useCartStore((s) => s.items)
  const remove = useCartStore((s) => s.remove)

  if (!items.length) return <div className="text-gray-300">Your cart is empty</div>

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4 text-white">Cart</h1>
      <div className="bg-black/60 rounded-lg shadow-neon p-4 border border-white/5">
        <ul>
          {items.map((it) => (
            <li key={it.id} className="flex items-center justify-between py-3 border-b last:border-b-0 border-white/5">
              <div className="flex items-center gap-4">
                <div className="w-16 h-12 bg-white/5 rounded flex items-center justify-center text-gray-400">Img</div>
                <div>
                  <div className="font-medium text-white">{it.name}</div>
                  <div className="text-sm text-gray-300">${it.price}</div>
                </div>
              </div>
              <div className="flex items-center gap-3">
                <button className="text-neon-pink" onClick={() => remove(it.id)}>Remove</button>
              </div>
            </li>
          ))}
        </ul>
        <div className="mt-4 text-right font-semibold text-neon-cyan">Total: ${items.reduce((s, i) => s + (i.price || 0), 0)}</div>
      </div>
    </section>
  )
}
