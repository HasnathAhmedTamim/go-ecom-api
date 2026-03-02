import React from 'react'
import LandingHero from '../components/LandingHero'
import Features from '../components/Features'

const Landing = () => {
  return (
    <div className="min-h-screen bg-black text-gray-200">
      <LandingHero />
      <Features />
      <section className="bg-black/20 py-12">
        <div className="container mx-auto px-4 text-center">
          <h2 className="text-2xl font-semibold mb-4 text-white">Ready to try?</h2>
          <p className="mb-6 text-gray-300 max-w-xl mx-auto">Start your free trial — no credit card required. Deploy a sandbox in seconds and import demo data.</p>
          <div className="flex flex-col sm:flex-row justify-center gap-4">
            <a href="/register" className="px-6 py-3 bg-neon-pink text-black rounded-md shadow-neon hover:brightness-95">Get Started</a>
            <a href="/login" className="px-6 py-3 border border-neon-cyan text-neon-cyan rounded-md">Sign In</a>
          </div>
        </div>
      </section>
    </div>
  )
}

export default Landing
