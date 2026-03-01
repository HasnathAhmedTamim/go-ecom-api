import useOrders from '../hooks/useOrders'

export default function Orders() {
  const { data, isLoading, error } = useOrders()

  if (isLoading) return <div>Loading orders...</div>
  if (error) return <div>Error loading orders</div>

  return (
    <section>
      <h1 className="text-2xl font-semibold mb-4">My Orders</h1>
      <ul>
        {data?.map((o) => (
          <li key={o.id} className="py-2 border-b">
            <div className="font-medium">Order {o.id}</div>
            <div className="text-sm text-gray-600">Status: {o.status}</div>
            {o.total !== undefined && (
              <div className="text-sm text-gray-800">Total: ${Number(o.total).toFixed(2)}</div>
            )}
            {o.created_at && (
              <div className="text-xs text-gray-500">Placed: {new Date(o.created_at).toLocaleString()}</div>
            )}
          </li>
        ))}
      </ul>
    </section>
  )
}
