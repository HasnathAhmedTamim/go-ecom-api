
import './App.css'
import { useEffect, useState } from 'react'
import api from './api'

function App() {
  const [message, setMessage] = useState(null)
  const [loading, setLoading] = useState(true)
  const [err, setErr] = useState(null)

  useEffect(() => {
    let mounted = true
    api
      .get('/api/')
      .then((res) => {
        if (mounted) setMessage(res.data?.message || JSON.stringify(res.data))
      })
      .catch((e) => {
        if (mounted) setErr(e.message || 'request failed')
      })
      .finally(() => {
        if (mounted) setLoading(false)
      })

    return () => {
      mounted = false
    }
  }, [])

  return (
    <>
      <div className='text-gray-500 text-center mt-10'>
        {loading && 'Connecting to backend...'}
        {!loading && err && `Error: ${err}`}
        {!loading && !err && message}
      </div>
    </>
  )
}

export default App
