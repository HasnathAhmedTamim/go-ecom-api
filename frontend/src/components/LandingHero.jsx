import React from 'react'
import HeroMock from '../assets/hero-mock.svg'

const LandingHero = () => {
  return (
    <header role="banner" className="bg-black text-neon-cyan py-20">
      <div className="container mx-auto px-4 flex flex-col-reverse md:flex-row items-center gap-8">
        <div className="w-full md:w-1/2 text-center md:text-left">
          <h1 className="text-3xl sm:text-4xl md:text-5xl font-extrabold leading-tight text-white drop-shadow-neon">GameHub — Gear for Champions</h1>
          <p className="mt-4 text-gray-300 max-w-xl mx-auto md:mx-0">High-performance peripherals, ergonomics-first chairs, and pro-grade accessories with neon flair. Tune your setup for comfort and competitive advantage.</p>
          <div className="mt-6 flex justify-center md:justify-start gap-3">
            <a href="/products" aria-label="Shop now" className="px-5 py-3 bg-neon-pink text-black rounded-md font-semibold shadow-neon">Shop Now</a>
            <a href="/admin/products" aria-label="Admin" className="px-5 py-3 border border-neon-cyan text-neon-cyan rounded-md">Admin</a>
          </div>
          <ul className="mt-6 grid grid-cols-3 gap-3 text-gray-400 max-w-sm mx-auto md:mx-0">
            <li className="text-sm"><strong className="text-white">Free shipping</strong></li>
            <li className="text-sm"><strong className="text-white">30-day returns</strong></li>
            <li className="text-sm"><strong className="text-white">1-year warranty</strong></li>
          </ul>
        </div>

        <div className="w-full md:w-1/2 flex justify-center md:justify-end">
          <div className="bg-gradient-to-br from-white/5 to-white/2 rounded-xl p-4 hero-card shadow-neon w-full max-w-lg">
            <div className="rounded-lg overflow-hidden border border-white/5">
              <img src={HeroMock} alt="dashboard preview" className="w-full h-60 sm:h-72 md:h-80 object-cover" />
            </div>
            <div className="mt-3 text-sm text-gray-300">Featured: Aurora Wireless Gaming Headset — immersive 7.1 surround sound.</div>
          </div>
        </div>
      </div>
    </header>
  )
}

export default LandingHero
