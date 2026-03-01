export default function AdminDashboard() {
  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4">Admin Dashboard</h1>
      <div className="space-y-2">
        <a href="/admin/products" className="text-blue-600 block">Manage Products</a>
        <a href="/admin/orders" className="text-blue-600 block">Manage Orders</a>
      </div>
    </section>
  )
}
