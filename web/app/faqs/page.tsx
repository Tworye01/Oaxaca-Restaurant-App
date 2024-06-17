import React from 'react';
import { NavigationBarComponent } from '@/components/nav-bar';
import { AccordionComponent } from '@/components/accordion';

export default function FAQs() {
  const faqs = [
    {
      preview: 'What are your opening hours?',
      contents: 'We are open from 11 AM to 10 PM from Monday to Saturday, and from 11 AM to 9 PM on Sundays.',
    },
    {
      preview: 'Can I make a reservation?',
      contents: 'Yes, reservations are recommended, especially on weekends. Please call us to book a table.',
    },
    {
      preview: 'Is there a dress code?',
      contents: 'Our restaurant has a casual dress code. Come as you are and enjoy our authentic Mexican cuisine!',
    },
    {
      preview: 'Do you serve alcohol?',
      contents: 'Yes, we have a full bar with a selection of beers, wines, and signature cocktails.',
    },
    {
      preview: 'Is parking available at the restaurant?',
      contents: 'Yes, we have a parking lot available for our customers at no extra charge.',
    },
    {
      preview: 'Do you cater to allergies',
      contents: 'Yes, We take food allergies and dietary restrictions very seriously. Our kitchen staff is trained to handle allergy-related requests with the utmost care. Just inform the staff and we will cater to you.',
    },
  ];

  return (
    <main className="min-h-screen bg-warm-100">
      <NavigationBarComponent />

      <section className="py-14">
        <h1 className="text-4xl font-bold text-center mb-8">Frequently Asked Questions</h1>
        <div className="max-w-3xl mx-auto px-4">
          {
            faqs.map((faq, index) => (
              <AccordionComponent key={index} preview={<>{faq.preview}</>} contents={<>{faq.contents}</>} />
            ))
          }
        </div>
      </section>
    </main>
  );
};
