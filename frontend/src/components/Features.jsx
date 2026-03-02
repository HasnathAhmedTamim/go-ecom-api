import React from 'react'

const FeatureCard = ({title, children}) => (
  <div className="p-6 bg-black/40 rounded-lg shadow-sm border border-neon-cyan/10">
    <h3 className="text-lg font-semibold mb-2 text-white">{title}</h3>
    <p className="text-gray-300 text-sm">{children}</p>
  </div>
)

const Features = () => {
  return (
    <section className="py-12 bg-black">
      <div className="container mx-auto px-4">
        <h2 className="text-2xl font-semibold text-center mb-8 text-white">Why shop with GameHub</h2>
        <div className="feature-grid grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3">
          <FeatureCard title="Pro-grade Peripherals">Hand-picked keyboards, mice, headsets and controllers from top manufacturers.</FeatureCard>
          <FeatureCard title="Performance & Comfort">Ergonomic chairs and accessories designed for long sessions and competitive play.</FeatureCard>
          <FeatureCard title="Fast Shipping & Support">Fast dispatch, easy returns, and dedicated support for gamers.</FeatureCard>
        </div>
      </div>
    </section>
  )
}

export default Features
