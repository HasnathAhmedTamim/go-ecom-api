
import './App.css'
import { Routes, Route } from 'react-router-dom'
import Header from './components/Header'
import Footer from './components/Footer'
import Toaster from './components/Toaster'
import Home from './pages/Home'
import Products from './pages/Products'
import Product from './pages/Product'
import Cart from './pages/Cart'
import Login from './pages/Login'
import Checkout from './pages/Checkout'
import Orders from './pages/Orders'
import OrderConfirmation from './pages/OrderConfirmation'
import AdminRoute from './components/AdminRoute'
import AdminDashboard from './pages/AdminDashboard'
import PrivateRoute from './components/PrivateRoute'
import AdminProducts from './pages/AdminProducts'
import AdminOrders from './pages/AdminOrders'
import Register from './pages/Register'
import Landing from './pages/Landing'

function App() {
  return (
    <>
      <Header />
      <main className="container mx-auto p-4">
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route path="/home" element={<Home />} />
          <Route path="/products" element={<Products />} />
          <Route path="/products/:id" element={<Product />} />
          <Route path="/cart" element={<PrivateRoute><Cart /></PrivateRoute>} />
          <Route path="/checkout" element={<PrivateRoute><Checkout/></PrivateRoute>} />
          <Route path="/orders" element={<PrivateRoute><Orders/></PrivateRoute>} />
          <Route path="/order/:id" element={<PrivateRoute><OrderConfirmation/></PrivateRoute>} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register/>} />
          <Route path="/admin" element={<AdminRoute><AdminDashboard/></AdminRoute>} />
          <Route path="/admin/products" element={<AdminRoute><AdminProducts/></AdminRoute>} />
          <Route path="/admin/orders" element={<AdminRoute><AdminOrders/></AdminRoute>} />
        </Routes>
      </main>
      <Footer />
      <Toaster />
    </>
  )
}

export default App
