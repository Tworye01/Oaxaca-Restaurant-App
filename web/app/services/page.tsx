import React from 'react';
import { NavigationBarComponent } from '@/components/nav-bar';

export default function Services() {
  // Define the array of services provided at the Mexican restaurant
  const services = [
    {
      title: "Authentic Mexican Cuisine",
      description: "Indulge in the rich flavors of Mexico with our authentic dishes made from traditional recipes passed down through generations."
    },
    {
      title: "Catering Services",
      description: "Host your events with ease by utilizing our catering services. We offer a wide range of Mexican delicacies to suit any occasion."
    },
    {
      title: "Happy Hour Specials",
      description: "Join us for happy hour and enjoy special discounts on select drinks and appetizers. Relax and unwind with friends and family in a lively atmosphere."
    },
    // Add more services as needed
  ];

  // Render the services page
  return (
    <main className="text-center min-h-screen text-brown-dark-100 bg-warm-100">
      <NavigationBarComponent />

      <div className="pt-20">
        <h1 className="text-3xl font-bold mb-4">Our Services</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {services.map((service, index: number) => (
            <div key={index} className="p-4 border border-gray-200 rounded-md shadow-md">
              <h2 className="text-xl font-semibold mb-2">{service.title}</h2>
              <p className="text-gray-700">{service.description}</p>
            </div>
          ))}
        </div>
      </div>
    </main>
  );
}
