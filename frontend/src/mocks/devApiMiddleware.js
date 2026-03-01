// Enhanced Vite dev middleware to mock /api responses when backend is unavailable
export function devApiMock() {
  // simple in-memory mock state
  const users = [
    { id: 'u1', name: 'Demo User', email: 'demo@example.com', password: 'password', isAdmin: false },
    { id: 'admin', name: 'Admin', email: 'admin@example.com', password: 'admin', isAdmin: true },
  ]

  const products = [
    { id: '1', name: 'Growth Widget', price: 49, description: 'Widget for growth', image: '/screenshots/hero-1.svg' },
    { id: '2', name: 'Engage Suite', price: 99, description: 'Suite to engage users', image: '/screenshots/feature-grid.svg' },
  ]

  const orders = []

  function json(res, obj, code = 200) {
    res.statusCode = code
    res.setHeader('Content-Type', 'application/json')
    res.end(JSON.stringify(obj))
  }

  return {
    name: 'dev-api-mock',
    configureServer(server) {
      server.middlewares.use((req, res, next) => {
        if (!req.url.startsWith('/api')) return next()

        // parse simple body for POST requests
        let body = ''
        req.on('data', (chunk) => (body += chunk))

        req.on('end', () => {
          const method = req.method || 'GET'
          const url = req.url.split('?')[0]

          // Auth: login
          if (method === 'POST' && url === '/api/auth/login') {
            try {
              const payload = JSON.parse(body || '{}')
              const user = users.find((u) => u.email === payload.email && u.password === payload.password)
              if (!user) return json(res, { error: 'Invalid credentials' }, 401)
              const token = user.isAdmin ? 'admin-token' : 'user-token-' + user.id
              return json(res, { user: { id: user.id, name: user.name, email: user.email, isAdmin: user.isAdmin }, token })
            } catch (e) {
              return json(res, { error: 'Bad request' }, 400)
            }
          }

          // Auth: register
          if (method === 'POST' && url === '/api/auth/register') {
            try {
              const payload = JSON.parse(body || '{}')
              const id = 'u' + (users.length + 1)
              const newUser = { id, name: payload.name, email: payload.email, password: payload.password, isAdmin: false }
              users.push(newUser)
              return json(res, { id: newUser.id, name: newUser.name, email: newUser.email })
            } catch (e) {
              return json(res, { error: 'Bad request' }, 400)
            }
          }

          // Products list
          if (method === 'GET' && (url === '/api/products' || req.url.startsWith('/api/products?'))) {
            return json(res, { items: products })
          }

          // Product detail
          if (method === 'GET' && url.startsWith('/api/products/')) {
            const id = url.split('/').pop()
            const p = products.find((x) => x.id === id) || { id, name: `Product ${id}`, price: 59 }
            return json(res, p)
          }

          // Create order (user)
          if (method === 'POST' && url === '/api/user/orders') {
            try {
              const payload = JSON.parse(body || '{}')
              // get user from Authorization header
              const auth = (req.headers.authorization || '')
              const userId = auth.includes('user-token-') ? auth.split('user-token-')[1] : auth === 'admin-token' ? 'admin' : 'u1'
              const id = 'o' + (orders.length + 1)
              const order = { id, user_id: userId, products: payload.products || {}, status: 'created' }
              orders.push(order)
              return json(res, order, 201)
            } catch (e) {
              return json(res, { error: 'Bad request' }, 400)
            }
          }

          // Get user orders
          if (method === 'GET' && url === '/api/user/orders') {
            const auth = (req.headers.authorization || '')
            const userId = auth.includes('user-token-') ? auth.split('user-token-')[1] : auth === 'admin-token' ? 'admin' : 'u1'
            const userOrders = orders.filter((o) => o.user_id === userId)
            return json(res, userOrders)
          }

          // Admin: products
          if (url === '/api/admin/products') {
            if (method === 'GET') return json(res, products)
            if (method === 'POST') {
              try {
                const payload = JSON.parse(body || '{}')
                const id = String(products.length + 1)
                const p = { id, ...payload }
                products.push(p)
                return json(res, p, 201)
              } catch (e) {
                return json(res, { error: 'Bad request' }, 400)
              }
            }
          }

          // Admin: delete product
          if (method === 'DELETE' && url.startsWith('/api/admin/products/')) {
            const id = url.split('/').pop()
            const idx = products.findIndex((x) => x.id === id)
            if (idx >= 0) products.splice(idx, 1)
            return json(res, { success: true })
          }

          // Admin orders
          if (method === 'GET' && url === '/api/admin/orders') {
            return json(res, orders)
          }

          // fallback
          return json(res, { message: 'Mock: backend unavailable' })
        })
      })
    },
  }
}
