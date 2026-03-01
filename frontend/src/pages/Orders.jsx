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
          </li>
        ))}
      </ul>
    </section>
  )
}
