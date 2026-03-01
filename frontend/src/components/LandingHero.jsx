import React from 'react'
import HeroMock from '../assets/hero-mock.svg'

const LandingHero = () => {
  return (
    <header role="banner" className="bg-gradient-to-r from-indigo-600 to-indigo-400 text-white py-16 md:py-24">
      <div className="container mx-auto px-4 flex flex-col-reverse md:flex-row items-center gap-8">
        <div className="w-full md:w-1/2 text-center md:text-left">
          <h1 className="text-3xl sm:text-4xl md:text-5xl font-bold leading-tight">Analytics built for growth</h1>
          <p className="mt-4 text-indigo-100 max-w-xl mx-auto md:mx-0">A modular, fast dashboard with reusable widgets, easy billing, and granular permissions — ship faster and understand users in minutes.</p>
          <div className="mt-6 flex justify-center md:justify-start gap-3">
            <a href="/register" aria-label="Start free trial" className="px-5 py-3 bg-white text-indigo-600 rounded-md font-medium shadow">Start free trial</a>
            <a href="/login" aria-label="Sign in" className="px-5 py-3 border border-white text-white rounded-md">Sign in</a>
          </div>
          <ul className="mt-6 grid grid-cols-3 gap-3 text-indigo-100 max-w-sm mx-auto md:mx-0">
            <li className="text-sm"><strong>SSO-ready</strong></li>
            <li className="text-sm"><strong>Realtime</strong></li>
            <li className="text-sm"><strong>Multi-tenant</strong></li>
          </ul>
        </div>
        <div className="w-full md:w-1/2 flex justify-center md:justify-end">
          <div className="bg-white/10 rounded-xl p-3 hero-card shadow-lg w-full max-w-md">
            <div className="rounded-lg overflow-hidden">
              <img src={HeroMock} alt="dashboard preview" className="w-full h-44 sm:h-56 md:h-64 object-cover" />
            </div>
          </div>
        </div>
      </div>
    </header>
  )
}

export default LandingHero
