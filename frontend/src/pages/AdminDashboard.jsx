import OrderStats from '../components/OrderStats'

export default function AdminDashboard() {
  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4">Admin Dashboard</h1>
      <div className="grid grid-cols-2 gap-4">
        <div>
          <OrderStats />
        </div>
        <div className="p-4 border rounded bg-black/60 text-gray-300 border-white/5">
          <h2 className="font-medium text-white">Quick Links</h2>
          <div className="mt-2 space-y-2">
            <a href="/admin/products" className="text-neon-cyan block">Manage Products</a>
            <a href="/admin/orders" className="text-neon-cyan block">Manage Orders</a>
          </div>
        </div>
      </div>
    </section>
  )
}
