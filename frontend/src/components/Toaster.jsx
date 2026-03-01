import { useEffect } from 'react'
import useToastStore from '../stores/useToastStore'

function Toast({ t }) {
  const remove = useToastStore((s) => s.remove)
  useEffect(() => {
    const tm = setTimeout(() => remove(t.id), t.duration || 4000)
    return () => clearTimeout(tm)
  }, [t.id])

  const bg = t.type === 'error' ? 'bg-red-600' : t.type === 'success' ? 'bg-green-600' : 'bg-gray-800'

  return (
    <div className={`${bg} text-white px-4 py-2 rounded shadow-sm mb-2`}>{t.title || t.message}</div>
  )
}

export default function Toaster() {
  const toasts = useToastStore((s) => s.toasts)
  return (
    <div className="fixed right-4 top-4 z-50 flex flex-col items-end">
      {toasts.map((t) => (
        <Toast key={t.id} t={t} />
      ))}
    </div>
  )
}
