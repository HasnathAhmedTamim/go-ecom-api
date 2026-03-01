import React from 'react'

const FeatureCard = ({title, children}) => (
  <div className="p-6 bg-white rounded-lg shadow-sm">
    <h3 className="text-lg font-semibold mb-2">{title}</h3>
    <p className="text-gray-600 text-sm">{children}</p>
  </div>
)

const Features = () => {
  return (
    <section className="py-12">
      <div className="container mx-auto px-4">
        <h2 className="text-2xl font-semibold text-center mb-8">Built for product-led teams</h2>
        <div className="feature-grid grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3">
          <FeatureCard title="Modular Widgets">Drag-and-drop widgets with real-time updates and exportable reports.</FeatureCard>
          <FeatureCard title="Usage-based Billing">Flexible billing with usage tiers, invoices, and coupons out-of-the-box.</FeatureCard>
          <FeatureCard title="Permissions & Teams">Invite teammates, set roles, and audit actions easily.</FeatureCard>
        </div>
      </div>
    </section>
  )
}

export default Features
